package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	m "task-process-service/internal/monitoring"

	"task-process-service/internal/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(ser service.TaskService) TaskHandler {
	return TaskHandler{service: ser}
}

// @Summary Add a new task
// @Description Adds a new task to the task list
// @Tags Tasks
// @ID add-task
// @Accept  json
// @Produce  json
// @Param task body domain.TaskAddReq true "Task Add to Database"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
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

	id, err := h.service.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.TaskCreate.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary task Update
// @Description Update a task
// @Tags Tasks
// @ID update-task
// @Accept  json
// @Produce  json
// @Param task body domain.TaskUpdateReq true "Task object"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/update [put]
// @Security ApiKeyAuth
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

	id, err := h.service.Update(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.TaskUpdate.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Delete a task
// @Description Delete a task
// @Tags Tasks
// @ID delete-task
// @Accept  json
// @Produce  json
// @Param task body domain.TaskDeleteReq true "Task object"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/delete [delete]
// @Security ApiKeyAuth
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

	id, err := h.service.Delete(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.TaskDelete.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags Tasks
// @ID get-task
// @Accept  json
// @Produce  json
// @Param task body domain.TaskGetByIdReq true "Task object"
// @Success 200 {object} domain.TaskGetDataList
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/getById [get]
// @Security ApiKeyAuth
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

	task, err := h.service.GetById(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.TaskGetById.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// @Summary Get all tasks
// @Description Get all tasks
// @Tags Tasks
// @ID get-all-tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.TaskGetDataList
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/getAll [get]
// @Security ApiKeyAuth
func (h TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.TaskGetAll.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
