package domain

type TaskAddReq struct {
	Title        string
	Description  string
	StatusId     int
	CreateUserId int
}

type TaskUpdateReq struct {
	Id           int
	Title        string
	Description  string
	StatusId     int
	UpdateUserId int
}

type TaskDeleteReq struct {
	Id           int
	UpdateUserId int
}

type TaskGetByIdReq struct {
	Id int
}

type TaskGetResp struct {
	Status string
	Result UserGetDataList
}

type TaskGetAllResp struct {
	Status string
	Result []TaskGetDataList
}

type TaskGetDataList struct {
	Id          *int
	Title       *string
	Description *string
	StatusId    *int
	Status      *string
	Duration    int
	Difficulty  int
	CreatedAt   string
}
