package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	"task-process-service/internal/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) TaskHandler {
	return TaskHandler{service: service}
}

// @Summary Create Task
// @Description Create a new task
// @Tags task-create
// @ID create-task
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /tasks/add [post]
func (h TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.TaskAddReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.AddTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Update Task
// @Description Update a task
// @Tags task-update
// @ID update-task
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /tasks/update [put]
func (h TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.TaskUpdateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.UpdateTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Delete Task
// @Description Delete a task
// @Tags task-delete
// @ID delete-task
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /tasks/delete [delete]
func (h TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.TaskDeleteReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.DeleteTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Get Task by ID
// @Description Get a task by ID
// @Tags task-getById
// @ID get-task-by-id
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /tasks/getById [get]
func (h TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.TaskGetByIdReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetTaskById(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Get All Tasks
// @Description Get all tasks
// @Tags task-getAll
// @ID get-all-tasks
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /tasks/getAll [get]
func (h TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
