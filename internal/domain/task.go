package domain

type TaskAddReq struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	StatusId     int    `json:"statusId"`
	Difficulty   int    `json:"difficulty"`
	Duration     int    `json:"duration"`
	CreateUserId int    `json:"createUserId"`
}

type TaskUpdateReq struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	StatusId     int    `json:"statusId"`
	Difficulty   int    `json:"difficulty"`
	Duration     int    `json:"duration"`
	UpdateUserId int    `json:"updateUserId"`
}

type TaskDeleteReq struct {
	Id           int `json:"id"`
	UpdateUserId int `json:"updateUserId"`
}

type TaskGetByIdReq struct {
	Id int `json:"id"`
}

// swagger:response taskGetResponse
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

type TaskStatusGetDataList struct {
	Id   int
	Name string
}
