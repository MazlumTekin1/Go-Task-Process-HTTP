package handler

import (
	"encoding/json"
	"net/http"
	"task-process-service/internal/domain"
	"task-process-service/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(ser service.UserService) UserHandler {
	return UserHandler{service: ser}
}

// @Summary Create User
// @Description Create a new user
// @Tags user-create
// @ID create-user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /users/add [post]
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

	id, err := h.service.AddUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Update User
// @Description Update a user
// @Tags user-update
// @ID update-user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /users/update [put]
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

	id, err := h.service.UpdateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Delete User
// @Description Delete a user
// @Tags user-delete
// @ID delete-user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /users/delete [delete]
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

	id, err := h.service.DeleteUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// @Summary Get User By Id
// @Description Get a user by id
// @Tags user-get
// @ID get-user-by-id
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /users/getById [get]
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

	user, err := h.service.GetUserById(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Get All Users
// @Description Get all users
// @Tags user-getAll
// @ID get-all-users
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200
// @Router /users/getAll [get]
func (h UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.service.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
