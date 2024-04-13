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

func TestTaskHandler_Create(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("Create", mock.Anything).Return(1, nil)

	taskHandler := NewTaskHandler(mockService)

	task := domain.TaskAddReq{
		Title:        "Sample Task",
		Description:  "This is a sample task",
		Duration:     10,
		CreateUserId: 318836,
	}

	req := createTaskRequest(task)

	res := httptest.NewRecorder()

	taskHandler.Create(res, req)

	assertResponseCode(t, res.Code, http.StatusCreated)
	assertTaskCreation(t, res.Body.String())
	mockService.AssertCalled(t, "Create", task)
}

func createTaskRequest(task domain.TaskAddReq) *http.Request {
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks/add", bytes.NewBuffer(taskJSON))
	return req
}

func assertResponseCode(t *testing.T, actual, expected int) {
	t.Helper()
	assert.Equal(t, expected, actual, "expected response code %d, got %d", expected, actual)
}

func assertTaskCreation(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestTaskHandler_Update(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("Update", mock.Anything).Return(1, nil)

	taskHandler := NewTaskHandler(mockService)

	task := domain.TaskUpdateReq{
		Id:           1,
		Title:        "Sample Task",
		Description:  "This is a sample task",
		Duration:     10,
		UpdateUserId: 318836,
	}

	req := updateTaskRequest(task)

	res := httptest.NewRecorder()

	taskHandler.Update(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertTaskUpdate(t, res.Body.String())
	mockService.AssertCalled(t, "Update", task)
}

func updateTaskRequest(task domain.TaskUpdateReq) *http.Request {
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks/update", bytes.NewBuffer(taskJSON))
	return req
}

func assertTaskUpdate(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestTaskHandler_Delete(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("Delete", mock.Anything).Return(1, nil)

	taskHandler := NewTaskHandler(mockService)

	task := domain.TaskDeleteReq{
		Id:           1,
		UpdateUserId: 318836,
	}

	req := deleteTaskRequest(task)

	res := httptest.NewRecorder()

	taskHandler.Delete(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertTaskDelete(t, res.Body.String())
	mockService.AssertCalled(t, "Delete", task)
}

func deleteTaskRequest(task domain.TaskDeleteReq) *http.Request {
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("DELETE", "/tasks/delete", bytes.NewBuffer(taskJSON))
	return req
}

func assertTaskDelete(t *testing.T, responseBody string) {
	t.Helper()
	var response map[string]int
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotZero(t, response["id"])
}

func TestTaskHandler_GetById(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("GetById", mock.Anything).Return(domain.TaskGetDataList{}, nil)

	taskHandler := NewTaskHandler(mockService)

	task := domain.TaskGetByIdReq{
		Id: 1,
	}

	req := getByIdRequest(task)

	res := httptest.NewRecorder()

	taskHandler.GetById(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertTaskGetById(t, res.Body.String())
	mockService.AssertCalled(t, "GetById", task)
}

func getByIdRequest(task domain.TaskGetByIdReq) *http.Request {
	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("GET", "/tasks/getById", bytes.NewBuffer(taskJSON))
	return req
}

func assertTaskGetById(t *testing.T, responseBody string) {
	t.Helper()
	var response domain.TaskGetDataList
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response)
}

func TestTaskHandler_GetAll(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("GetAll").Return([]domain.TaskGetDataList{}, nil)

	taskHandler := NewTaskHandler(mockService)

	req := getAllRequest()

	res := httptest.NewRecorder()

	taskHandler.GetAll(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertTaskGetAll(t, res.Body.String())
	mockService.AssertCalled(t, "GetAll")
}

func getAllRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/tasks/getAll", nil)
	return req
}

func assertTaskGetAll(t *testing.T, responseBody string) {
	t.Helper()
	var response []domain.TaskGetDataList
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response)
}

func TestTaskHandler_GetAllTaskStatus(t *testing.T) {
	mockService := &m.MockTaskService{}
	mockService.On("GetAllTaskStatus").Return([]domain.TaskStatusGetDataList{}, nil)

	taskHandler := NewTaskHandler(mockService)

	req := getAllTaskStatusRequest()

	res := httptest.NewRecorder()

	taskHandler.GetAllTaskStatus(res, req)

	assertResponseCode(t, res.Code, http.StatusOK)
	assertTaskGetAllTaskStatus(t, res.Body.String())
	mockService.AssertCalled(t, "GetAllTaskStatus")
}

func getAllTaskStatusRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/tasks/taskStatus", nil)
	return req
}

func assertTaskGetAllTaskStatus(t *testing.T, responseBody string) {
	t.Helper()
	var response []domain.TaskStatusGetDataList
	err := json.Unmarshal([]byte(responseBody), &response)
	assert.NoError(t, err)
	assert.NotZero(t, response)
}
