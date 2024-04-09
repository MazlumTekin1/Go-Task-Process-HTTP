package handler

import (
	"encoding/json"
	"net/http"

	"task-process-service/internal/domain"
	m "task-process-service/internal/monitoring"
	"task-process-service/internal/service"
)

type AuthHandler struct {
	AuthService service.AuthService
	//	JWTService  service.JWTService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{AuthService: service}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req domain.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resLogin, err := handler.AuthService.LoginService(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Login.Inc()
	w.Header().Set("Content-Type", "application/json")
	if resLogin.IsValid {
		json.NewEncoder(w).Encode(resLogin)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
