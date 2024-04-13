package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	m "task-process-service/internal/monitoring"
	"task-process-service/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(ser service.UserService) UserHandler {
	return UserHandler{service: ser}
}

// @Summary Add a new user
// @Description Adds a new user to the user list
// @Tags Users
// @ID add-user
// @Accept  json
// @Produce  json
// @Param user body domain.UserAddReq true "User Add to Database"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/add [post]
// @Security ApiKeyAuth
func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.UserAddReq
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
	m.UserCreate.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Update a user
// @Description Update a user
// @Tags Users
// @ID update-user
// @Accept  json
// @Produce  json
// @Param user body domain.UserUpdateReq true "User object"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/update [put]
// @Security ApiKeyAuth
func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.UserUpdateReq
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
	m.UserUpdate.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param user body domain.UserDeleteReq true "User object"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/delete [delete]
// @Security ApiKeyAuth
func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.UserDeleteReq
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
	m.UserDelete.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags Users
// @ID get-user
// @Accept  json
// @Produce  json
// @Param user body domain.UserGetByIdReq true "User object"
// @Success 200 {object} domain.UserGetDataList
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/getById [get]
// @Security ApiKeyAuth
func (h UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req domain.UserGetByIdReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetById(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.UserGetById.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Get all users
// @Description Get all users
// @Tags Users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.UserGetDataList
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/getAll [get]
// @Security ApiKeyAuth
func (h UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.UserGetAll.Inc()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
