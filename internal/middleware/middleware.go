package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"task-process-service/internal/service"
	"time"
)

type visitor struct {
	lastSeen time.Time
	visits   int
}

var visitors = make(map[string]*visitor)
var mtx sync.Mutex

type MiddlewareInterface interface {
	AuthMiddleware(next http.HandlerFunc) http.HandlerFunc
	RateLimit(next http.HandlerFunc) http.HandlerFunc
}

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

func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtx.Lock()
		defer mtx.Unlock()

		visitorIP := r.RemoteAddr
		v, exists := visitors[visitorIP]
		if !exists {
			visitors[visitorIP] = &visitor{time.Now(), 1}
		} else if time.Since(v.lastSeen) > 1*time.Minute {
			v.lastSeen = time.Now()
			v.visits = 1
		} else {
			v.visits++
			if v.visits > 10 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(map[string]string{"error": "You have reached maximum request limit."})
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}
