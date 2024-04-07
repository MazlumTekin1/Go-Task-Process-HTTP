package service

import (
	"fmt"
	"task-process-service/internal/repository"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secretKey string
	repo      repository.UserRepository
}

func NewJWTAuthService() *JWTService {
	return &JWTService{
		secretKey: "simdi-secret-key-belirleme-zamani:)",
		repo:      repository.NewUserRepository(nil),
	}
}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		encodedToken,
		&authCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.secretKey), nil
		},
	)
}
