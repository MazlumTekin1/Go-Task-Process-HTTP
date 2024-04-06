package service

/*
type JWTService interface {
	GenerateToken(email string, id int) string
	ValidateToken(token string) (*jwt.Token, error)
}
*/

// type authCustomClaims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

// type jwtServices struct {
// 	secretKey string
// 	issure    string
// }

// func NewJWTAuthService() JWTService {
// 	return &jwtServices{
// 		secretKey: "secretKey",
// 		issure:    "test.com",
// 	}
// }

// func (s *jwtServices) GenerateToken(email string, id int) string {
// 	claims := &authCustomClaims{
// 		email,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
// 			Issuer:    s.issure,
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	newToken, err := token.SignedString([]byte(s.secretKey))
// 	if err != nil {
// 		panic(err)
// 	}

// 	return newToken
// }

// func (s *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
// 	return jwt.ParseWithClaims(
// 		encodedToken,
// 		&authCustomClaims{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte(s.secretKey), nil
// 		},
// 	)
// }
