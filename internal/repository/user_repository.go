package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"task-process-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(User domain.UserAddReq) (int, error)
	Update(User domain.UserUpdateReq) (int, error)
	Delete(User domain.UserDeleteReq) (int, error)
	GetById(User domain.UserGetByIdReq) (domain.UserGetDataList, error)
	GetAll() ([]domain.UserGetDataList, error)
}

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (pg *UserRepositoryImpl) Create(User domain.UserAddReq) (int, error) {
	createDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	h := sha256.New()
	_, err := h.Write([]byte(User.Password))
	if err != nil {
		log.Println("user_repository.go -> UserCreate -> hash error: ", err)
	}
	passwordHash := hex.EncodeToString(h.Sum(nil))
	qInsert := `insert into test.users (first_name, last_name, email, password, created_at, create_user_id)
    values($1, $2, $3, $4, $5, $6)
	returning id;`
	err = pg.db.QueryRow(context.Background(), qInsert,
		User.FirstName,
		User.LastName,
		User.Email,
		passwordHash,
		&createDate,
		User.CreateUserId,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("UserRepositoryImpl.Create", err)
	}

	return id, nil
}

func (pg *UserRepositoryImpl) Update(User domain.UserUpdateReq) (int, error) {
	updateDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	qInsert := `UPDATE test.users
	SET first_name=$1, last_name=$2, email=$3,  
	updated_at=$4, update_user_id=$5
	where id = $6 and is_deleted = false
	returning id;`
	err := pg.db.QueryRow(context.Background(), qInsert,
		User.FirstName,
		User.LastName,
		User.Email,
		&updateDate,
		User.UpdateUserId,
		User.Id,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("UserRepositoryImpl.Update", err)
	}

	return id, nil
}

func (pg *UserRepositoryImpl) Delete(User domain.UserDeleteReq) (int, error) {
	DeleteDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	qInsert := `update test.users 
	set is_deleted = true, updated_at = $1, update_user_id = $2
	where id = $3 and is_deleted = false
	returning id;`
	err := pg.db.QueryRow(context.Background(), qInsert,
		&DeleteDate,
		User.UpdateUserId,
		User.Id,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("UserRepositoryImpl.Delete", err)
	}

	return id, nil
}

func (pg *UserRepositoryImpl) GetById(User domain.UserGetByIdReq) (domain.UserGetDataList, error) {
	res := domain.UserGetDataList{}
	var createdAt time.Time
	qInsert := `SELECT 
	id, first_name, last_name, email, created_at, user_level
	FROM test.users
	where is_deleted = false and id = $1`
	err := pg.db.QueryRow(context.Background(), qInsert, User.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&createdAt,
		&res.DeveloperWorkHourDifficulty,
	)
	if err != nil {
		return res, domain.NewDatabaseError("UserRepositoryImpl.GetById", err)
	}
	res.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	return res, nil
}

func (pg *UserRepositoryImpl) GetAll() ([]domain.UserGetDataList, error) {

	res := []domain.UserGetDataList{}
	var createdAt time.Time
	qInsert := `SELECT 
	id, first_name, last_name, email, created_at, user_level
	FROM test.users
	where is_deleted = false
	order by id desc;`
	rows, err := pg.db.Query(context.Background(), qInsert)

	if err != nil {
		return res, domain.NewDatabaseError("UserRepositoryImpl.GetAll", err)
	}
	defer rows.Close()

	for rows.Next() {
		row := domain.UserGetDataList{}
		err = rows.Scan(
			&row.Id,
			&row.FirstName,
			&row.LastName,
			&row.Email,
			&createdAt,
			&row.DeveloperWorkHourDifficulty,
		)
		if err != nil {
			return res, domain.NewDatabaseError("UserRepositoryImpl.GetAll", err)
		}
		row.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		res = append(res, row)
	}
	return res, nil
}
