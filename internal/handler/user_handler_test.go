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
)

func TestUserHandler_Create(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)

	user := domain.UserAddReq{FirstName: "Test User", LastName: "Test Last Name", Email: "test@test.com", Password: "Test Password", CreateUserId: 1}
	mockRepo.On("Create", user).Return(1, nil)

	handler := NewUserHandler(mockRepo)

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/users/add", bytes.NewBuffer(body))
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

func TestUserHandler_Update(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)

	user := domain.UserUpdateReq{Id: 1, FirstName: "Test User", LastName: "Test Last Name", Email: "test.user@test.com", Password: "Test Password", UpdateUserId: 1}
	mockRepo.On("Update", user).Return(1, nil)

	handler := NewUserHandler(mockRepo)

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("PUT", "/users/update", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected, actual map[string]interface{}
	json.Unmarshal([]byte(`{"id":1}`), &expected)
	json.Unmarshal(rr.Body.Bytes(), &actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestUserDelete(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)

	user := domain.UserDeleteReq{Id: 1, DeleteUserId: 1}
	mockRepo.On("Delete", user).Return(1, nil)

	handler := NewUserHandler(mockRepo)

	body, _ := json.Marshal(user)
	req, err := http.NewRequest("DELETE", "/users/delete", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.Delete(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected, actual map[string]interface{}
	json.Unmarshal([]byte(`{"id":1}`), &expected)
	json.Unmarshal(rr.Body.Bytes(), &actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}
