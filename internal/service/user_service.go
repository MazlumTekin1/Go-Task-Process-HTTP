package service

import "task-process-service/internal/domain"

type UserService interface {
	AddUser(user domain.User) error
	DeleteUser(id int) error
	UpdateUser(user domain.User) error
	ListUsers() ([]domain.User, error)
}
