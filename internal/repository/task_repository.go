package repository

import (
	"context"
	"time"

	"task-process-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository interface {
	Create(task domain.TaskAddReq) (int, error)
	Update(task domain.TaskUpdateReq) (int, error)
	Delete(task domain.TaskDeleteReq) (int, error)
	GetById(task domain.TaskGetByIdReq) (domain.TaskGetDataList, error)
	GetAll() ([]domain.TaskGetDataList, error)
}

type TaskRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (pg *TaskRepositoryImpl) Create(task domain.TaskAddReq) (int, error) {
	createDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	qInsert := `insert into test.tasks (title, description, status_id, created_at, create_user_id)
    values($1, $2, $3, $4, $5)
	returning id;`
	err := pg.db.QueryRow(context.Background(), qInsert,
		task.Title,
		task.Description,
		task.StatusId,
		&createDate,
		task.CreateUserId,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("TaskRepositoryImpl.Create", err)
	}

	return id, nil
}

func (pg *TaskRepositoryImpl) Update(task domain.TaskUpdateReq) (int, error) {
	updateDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	qInsert := `UPDATE test.tasks
	SET title=$1, description=$2, status_id=$3, updated_at=$4, 
	update_user_id=$5
	where id = $6
	returning id;`
	err := pg.db.QueryRow(context.Background(), qInsert,
		task.Title,
		task.Description,
		task.StatusId,
		&updateDate,
		task.UpdateUserId,
		task.Id,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("TaskRepositoryImpl.Update", err)
	}

	return id, nil
}

func (pg *TaskRepositoryImpl) Delete(task domain.TaskDeleteReq) (int, error) {
	DeleteDate := time.Now().Format("2006-01-02 15:04:05")
	var id int

	qInsert := `update test.tasks 
	set is_deleted = true, updated_at = $1, update_user_id = $2
	where id = $3 and is_deleted = false
	returning id;`
	err := pg.db.QueryRow(context.Background(), qInsert,
		&DeleteDate,
		task.UpdateUserId,
		task.Id,
	).Scan(&id)
	if err != nil {
		return 0, domain.NewDatabaseError("TaskRepositoryImpl.Delete", err)
	}

	return id, nil
}

func (pg *TaskRepositoryImpl) GetById(task domain.TaskGetByIdReq) (domain.TaskGetDataList, error) {
	res := domain.TaskGetDataList{}
	var createdAt time.Time
	qInsert := `SELECT 
	main.id, main.title, main.description, main.status_id, ts.task_statu, main.created_at, main.difficulty 
	FROM test.tasks main
	inner join test.task_status as ts
	on main.status_id = ts.id
	where main.is_deleted = false and main.id = $1`
	err := pg.db.QueryRow(context.Background(), qInsert, task.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.StatusId,
		&res.Status,
		&createdAt,
		&res.Difficulty,
	)
	if err != nil {
		return res, domain.NewDatabaseError("TaskRepositoryImpl.GetById", err)
	}
	res.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	return res, nil
}

func (pg *TaskRepositoryImpl) GetAll() ([]domain.TaskGetDataList, error) {

	res := []domain.TaskGetDataList{}
	var createdAt time.Time
	qInsert := `SELECT 
	main.id, main.title, main.description, main.status_id, ts.task_statu, main.created_at, main.difficulty 
	FROM test.tasks main
	inner join test.task_status as ts
	on main.status_id = ts.id
	where main.is_deleted = false
	order by id desc;`
	rows, err := pg.db.Query(context.Background(), qInsert)

	if err != nil {
		return res, domain.NewDatabaseError("TaskRepositoryImpl.GetAll", err)
	}
	defer rows.Close()

	for rows.Next() {
		row := domain.TaskGetDataList{}
		err = rows.Scan(
			&row.Id,
			&row.Title,
			&row.Description,
			&row.StatusId,
			&row.Status,
			&createdAt,
			&row.Difficulty,
		)
		if err != nil {
			return res, domain.NewDatabaseError("TaskRepositoryImpl.GetAll", err)
		}
		row.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		res = append(res, row)
	}
	return res, nil
}
