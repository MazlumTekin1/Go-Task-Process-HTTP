package handler

import (
	"encoding/json"
	"net/http"

	"task-process-service/internal/domain"
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
	if resLogin.IsValid {
		json.NewEncoder(w).Encode(resLogin)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
	/*
		if isValid {
			log.Println("User is valid, generating token")
			token := handler.AuthService.GenerateToken(req.Email, id)
			// token := strconv.Itoa(id)
			w.Write([]byte(token))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
		}
	*/
}
