package domain

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	IsValid    bool
	UserId     int
	Token      string
	ExpireTime string
}

type CheckUserResponse struct {
	IsValid  bool
	UserId   int
	Email    string
	Password string
}
