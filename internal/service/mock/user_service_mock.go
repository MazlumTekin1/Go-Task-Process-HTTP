package service

import (
	"task-process-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user domain.UserAddReq) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) GetAll() ([]domain.UserGetDataList, error) {
	args := m.Called()
	return args.Get(0).([]domain.UserGetDataList), args.Error(1)
}

func (m *MockUserService) GetById(user domain.UserGetByIdReq) (domain.UserGetDataList, error) {
	args := m.Called(user)
	return args.Get(0).(domain.UserGetDataList), args.Error(1)
}

func (m *MockUserService) Update(user domain.UserUpdateReq) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) Delete(user domain.UserDeleteReq) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}
