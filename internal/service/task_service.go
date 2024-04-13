package service

import (
	"strconv"
	"task-process-service/internal/domain"
	"task-process-service/internal/repository"
)

type TaskService interface {
	Create(req domain.TaskAddReq) (int, error)
	Update(req domain.TaskUpdateReq) (int, error)
	Delete(req domain.TaskDeleteReq) (int, error)
	GetById(req domain.TaskGetByIdReq) (domain.TaskGetDataList, error)
	GetAll() ([]domain.TaskGetDataList, error)
	GetAllTaskStatus() ([]domain.TaskStatusGetDataList, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: repo,
	}
}

func (s *taskService) Create(req domain.TaskAddReq) (int, error) {
	if req.CreateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Create", `CreateUserId has to be more than zero , your User Id is `+strconv.Itoa(req.CreateUserId)+``)
	}
	if req.Title == "" || req.Description == "" || req.StatusId == 0 || req.Duration == 0 || req.Difficulty == 0 {
		return 0, domain.NewTaskFieldIsRequiredError("Create", "Title, Description, and Status are required")
	}
	id, err := s.taskRepo.Create(req)
	return id, err
}

func (s *taskService) Update(req domain.TaskUpdateReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Update", `UpdateUserId has to be more than zero, your User Id is `+strconv.Itoa(req.UpdateUserId)+``)
	}
	return s.taskRepo.Update(req)
}

func (s *taskService) Delete(req domain.TaskDeleteReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Delete", `UpdateUserId has to be more than zero, your User Id is `+strconv.Itoa(req.UpdateUserId)+``)
	}
	return s.taskRepo.Delete(req)
}

func (s *taskService) GetById(req domain.TaskGetByIdReq) (domain.TaskGetDataList, error) {
	if req.Id <= 0 {
		return domain.TaskGetDataList{}, domain.NewIdLessThanZeroError("GetById", `UserId has to be more than zero, your User Id is `+strconv.Itoa(req.Id)+``)
	}
	return s.taskRepo.GetById(req)
}

func (s *taskService) GetAll() ([]domain.TaskGetDataList, error) {
	return s.taskRepo.GetAll()
}
func (s *taskService) GetAllTaskStatus() ([]domain.TaskStatusGetDataList, error) {
	return s.taskRepo.GetAllTaskStatus()
}
