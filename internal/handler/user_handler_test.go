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

func TestUserHandler_Create(t *testing.T) {
	mockService := &m.MockUserService{}
	mockService.On("Create", mock.Anything).Return(1, nil)

	userHandler := NewUserHandler(mockService)

	user := domain.UserAddReq{

		CreateUserId: 318836,
	}

	req := createUserRequest(user)

	res := httptest.NewRecorder()

	userHandler.Create(res, req)

	assertResponseCode(t, res.Code, http.StatusCreated)
	assertUserCreation(t, res.Body.String())
	mockService.AssertCalled(t, "Create", user)
}

func createUserRequest(user domain.UserAddReq) *http.Request {
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/add", bytes.NewBuffer(userJSON))
	return req
}

func assertUserCreation(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestUserHandler_Update(t *testing.T) {
	mockService := &m.MockUserService{}
	mockService.On("Update", mock.Anything).Return(1, nil)

	userHandler := NewUserHandler(mockService)

	user := domain.UserUpdateReq{
		Id:           1,
		UpdateUserId: 318836,
	}

	req := updateUserRequest(user)

	res := httptest.NewRecorder()

	userHandler.Update(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertUserUpdate(t, res.Body.String())
	mockService.AssertCalled(t, "Update", user)
}

func updateUserRequest(user domain.UserUpdateReq) *http.Request {
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("PUT", "/users/update", bytes.NewBuffer(userJSON))
	return req
}

func assertUserUpdate(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestUserHandler_Delete(t *testing.T) {
	mockService := &m.MockUserService{}
	mockService.On("Delete", mock.Anything).Return(1, nil)

	userHandler := NewUserHandler(mockService)

	user := domain.UserDeleteReq{
		Id:           1,
		UpdateUserId: 318836,
	}

	req := deleteUserRequest(user)

	res := httptest.NewRecorder()

	userHandler.Delete(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertUserDelete(t, res.Body.String())
	mockService.AssertCalled(t, "Delete", user)
}

func deleteUserRequest(user domain.UserDeleteReq) *http.Request {
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("DELETE", "/users/delete", bytes.NewBuffer(userJSON))
	return req
}

func assertUserDelete(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestUserHandler_GetById(t *testing.T) {
	mockService := &m.MockUserService{}
	mockService.On("GetById", mock.Anything).Return(domain.UserGetDataList{}, nil)

	userHandler := NewUserHandler(mockService)

	user := domain.UserGetByIdReq{
		Id: 1,
	}

	req := getByIdRequest(domain.TaskGetByIdReq(user))

	res := httptest.NewRecorder()

	userHandler.GetById(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertUserGetById(t, res.Body.String())
	mockService.AssertCalled(t, "GetById", user)
}

func assertUserGetById(t *testing.T, responseBody string) {
	t.Helper()
	var response domain.UserGetDataList
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response)
}

func TestUserHandler_GetAll(t *testing.T) {
	mockService := &m.MockUserService{}
	mockService.On("GetAll").Return([]domain.UserGetDataList{}, nil)

	userHandler := NewUserHandler(mockService)

	req := getAllRequest()

	res := httptest.NewRecorder()

	userHandler.GetAll(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertUserGetAll(t, res.Body.String())
	mockService.AssertCalled(t, "GetAll")
}

func assertUserGetAll(t *testing.T, responseBody string) {
	t.Helper()
	var response []domain.UserGetDataList
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response)
}
