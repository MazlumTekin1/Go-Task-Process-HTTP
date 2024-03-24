package repository

import (
	"task-process-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task domain.TaskAddReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskRepository) GetAll() ([]domain.TaskGetDataList, error) {
	args := m.Called()
	return args.Get(0).([]domain.TaskGetDataList), args.Error(1)
}

func (m *MockTaskRepository) GetById(task domain.TaskGetByIdReq) (domain.TaskGetDataList, error) {
	args := m.Called(task)
	return args.Get(0).(domain.TaskGetDataList), args.Error(1)
}

func (m *MockTaskRepository) Update(task domain.TaskUpdateReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *MockTaskRepository) Delete(task domain.TaskDeleteReq) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}
