package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-process-service/internal/domain"
	m "task-process-service/internal/service/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_Login(t *testing.T) {
	mockService := &m.MockAuthService{}
	mockService.On("Login", mock.Anything).Return(1, nil)

	authHandler := NewAuthHandler(mockService)

	auth := domain.LoginRequest{
		Email:    "admin@test.com",
		Password: "12345",
	}

	req := loginAuthRequest(auth)

	res := httptest.NewRecorder()

	authHandler.Login(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertAuthCreation(t, res.Body.String())
	mockService.AssertCalled(t, "LoginService", auth)

}

func loginAuthRequest(task domain.LoginRequest) *http.Request {
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("GET", "/login", bytes.NewBuffer(taskJSON))
	return req
}

func assertAuthCreation(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}
