package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"task-process-service/internal/domain"
	"task-process-service/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Email    string
	Password string
}

type AuthService interface {
	LoginService(req domain.LoginRequest) (domain.LoginResponse, error)
	GenerateToken(email string, id int) (string, int)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		authRepo: repo,
	}
}

func (s *authService) LoginService(req domain.LoginRequest) (domain.LoginResponse, error) {
	var res domain.LoginResponse
	checkUser, err := s.authRepo.Login(req)
	if err != nil {
		return res, err
	}

	res.UserId = checkUser.UserId

	hashPassword := s.GenerateHashPassword(req.Password)

	if checkUser.Password == hashPassword {
		token, expireTime := s.GenerateToken(req.Email, checkUser.UserId)
		res.IsValid = true
		res.Token = token
		res.ExpireTime = strconv.Itoa(expireTime) + " seconds"
		return res, nil
	} else {
		res.IsValid = false
		res.Token = ""
		return res, nil
	}
}

type authCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//type jwtServices struct {
//secretKey string
//	issure    string
//}

// func NewJWTAuthService() JWTService {
// 	return &jwtServices{
// 		secretKey: "secretKey",
// 		issure:    "test.com",
// 	}
// }

func (s *authService) GenerateHashPassword(password string) string {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		log.Println("auth.go -> Login -> hash error: ", err)
	}
	passwordHash := hex.EncodeToString(h.Sum(nil))
	return passwordHash
}

func (s *authService) GenerateToken(email string, id int) (string, int) {
	secretKey := "secretKey"
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return newToken, 3600
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := "secretKey"
	return jwt.ParseWithClaims(
		encodedToken,
		&authCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		},
	)
}
