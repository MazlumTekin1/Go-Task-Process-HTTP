package middleware

import (
	"net/http"
	"strings"
	"task-process-service/internal/service"
)

func AuthMiddleware(jwtService service.JWTService) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			// The Authorization header is expected to be "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
				return
			}

			token := parts[1]

			// Validate the token using the jwtService
			ValidatedToken, err := jwtService.ValidateToken(token)
			if err != nil || ValidatedToken == nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// If the token is valid, call the next handler
			next.ServeHTTP(w, r)
		}
	}
}

/*
func AuthMiddleware(jwtService service.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				bearerToken := strings.Split(authHeader, " ")
				if len(bearerToken) == 2 {
					token, err := jwtService.ValidateToken(bearerToken[1])
					if token.Valid {
						next.ServeHTTP(w, r)
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte(err.Error()))
						return
					}
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}
*/
