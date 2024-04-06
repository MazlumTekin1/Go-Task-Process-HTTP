package domain

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	IsValid    bool   `json:"isValid"`
	UserId     int    `json:"userId"`
	Token      string `json:"token"`
	ExpireTime string `json:"expireTime"`
}

type CheckUserResponse struct {
	IsValid  bool
	UserId   int
	Email    string
	Password string
}
