package service

import (
	"task-process-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) Create(task domain.TaskAddReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskService) GetAll() ([]domain.TaskGetDataList, error) {
	args := m.Called()
	return args.Get(0).([]domain.TaskGetDataList), args.Error(1)
}
func (m *MockTaskService) GetAllTaskStatus() ([]domain.TaskStatusGetDataList, error) {
	args := m.Called()
	return args.Get(0).([]domain.TaskStatusGetDataList), args.Error(1)
}

func (m *MockTaskService) GetById(task domain.TaskGetByIdReq) (domain.TaskGetDataList, error) {
	args := m.Called(task)
	return args.Get(0).(domain.TaskGetDataList), args.Error(1)
}

func (m *MockTaskService) Update(task domain.TaskUpdateReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskService) Delete(task domain.TaskDeleteReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}
