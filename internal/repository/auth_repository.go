package repository

import (
	"context"
	"errors"

	"task-process-service/internal/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	Login(req domain.LoginRequest) (domain.CheckUserResponse, error)
	GenerateToken(email string, id int) string
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (pg *AuthRepositoryImpl) Login(req domain.LoginRequest) (domain.CheckUserResponse, error) {
	var cUser domain.CheckUserResponse

	err := pg.db.QueryRow(context.Background(),
		"SELECT id, email, password FROM test.users WHERE email = $1", req.Email).Scan(
		&cUser.UserId,
		&cUser.Email,
		&cUser.Password,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return cUser, errors.New("user not found")
		} else {
			return cUser, err
		}
	}
	return cUser, nil
}

func (pg *AuthRepositoryImpl) GenerateToken(email string, id int) string {
	return ""
}

func (pg *AuthRepositoryImpl) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
