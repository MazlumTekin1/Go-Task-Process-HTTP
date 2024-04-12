package service

import (
	"task-process-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(auth domain.LoginRequest) (domain.LoginResponse, error) {
	args := m.Called(auth)
	return args.Get(0).(domain.LoginResponse), args.Error(1)
}

func (m *MockAuthService) GenerateToken(email string, id int) (string, int) {
	args := m.Called(email, id)
	return args.String(0), args.Int(1)
}
