package domain

type User struct {
	Id    int
	Name  string
	Tasks []Task
}

type UserAddReq struct {
	FirstName    string
	LastName     string
	Email        string
	Password     string
	CreateUserId int
}

type UserUpdateReq struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	Password     string
	UpdateUserId int
}

type UserDeleteReq struct {
	Id           int
	DeleteUserId int
}

type UserGetByIdReq struct {
	Id int
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
	Id        *int
	FirstName *string
	LastName  *string
	Email     *string
	CreatedAt string
}