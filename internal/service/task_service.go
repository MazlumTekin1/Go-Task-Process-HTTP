package service

import "task-process-service/internal/domain"

type TaskService interface {
	AddTask(task domain.Task) error
	DeleteTask(id int) error
	UpdateTask(task domain.Task) error
	ListTasks() ([]domain.Task, error)
}
