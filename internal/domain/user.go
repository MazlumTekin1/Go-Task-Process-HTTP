package domain

type UserAddReq struct {
	FirstName                   string `json:"firstName"`
	LastName                    string `json:"lastName"`
	Email                       string `json:"email"`
	Password                    string `json:"password"`
	DeveloperWorkHourDifficulty int    `json:"developerLevel"`
	CreateUserId                int    `json:"createUserId"`
}

type UserUpdateReq struct {
	Id                          int    `json:"id"`
	FirstName                   string `json:"firstName"`
	LastName                    string `json:"lastName"`
	Email                       string `json:"email"`
	DeveloperWorkHourDifficulty int    `json:"developerLevel"`
	UpdateUserId                int    `json:"updateUserId"`
}

type UserDeleteReq struct {
	Id           int `json:"id"`
	UpdateUserId int `json:"updateUserId"`
}

type UserGetByIdReq struct {
	Id int `json:"id"`
}

type UserGetResp struct {
	Status string
	Result UserGetDataList
}

type UserGetAllResp struct {
	Status string
	Result []UserGetDataList
}

type UserGetDataList struct {
	Id                          int
	FirstName                   string
	LastName                    string
	Email                       string
	Name                        string
	DeveloperWorkHourDifficulty int
	CreatedAt                   string
}
