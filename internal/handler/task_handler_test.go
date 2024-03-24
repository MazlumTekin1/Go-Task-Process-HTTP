package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"task-process-service/internal/domain"
	mock "task-process-service/internal/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_Create(t *testing.T) {
	mockRepo := new(mock.MockTaskRepository)

	task := domain.TaskAddReq{Title: "Test Task", Description: "Test Description", StatusId: 1, CreateUserId: 1}
	mockRepo.On("Create", task).Return(1, nil)

	handler := NewTaskHandler(mockRepo)

	body, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/tasks/add", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var expected, actual map[string]interface{}
	json.Unmarshal([]byte(`{"id":1}`), &expected)
	json.Unmarshal(rr.Body.Bytes(), &actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestTaskHandler_Update(t *testing.T) {
	mockRepo := new(mock.MockTaskRepository)

	task := domain.TaskUpdateReq{Id: 1, Title: "Test Task", Description: "Test Description", StatusId: 1, UpdateUserId: 1}
	mockRepo.On("Update", task).Return(1, nil)

	handler := NewTaskHandler(mockRepo)

	body, _ := json.Marshal(task)
	req, err := http.NewRequest("PUT", "/tasks/update", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `{"id":1}`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}

func TestTaskHandler_Delete(t *testing.T) {
	mockRepo := new(mock.MockTaskRepository)

	task := domain.TaskDeleteReq{Id: 1, UpdateUserId: 1}
	mockRepo.On("Delete", task).Return(1, nil)

	handler := NewTaskHandler(mockRepo)

	body, _ := json.Marshal(task)
	req, err := http.NewRequest("DELETE", "/tasks/delete", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("Hata olustu: ", err)
	}

	rr := httptest.NewRecorder()

	handler.Delete(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `{"id":1}`
	assert.JSONEq(t, expectedBody, rr.Body.String())
}
