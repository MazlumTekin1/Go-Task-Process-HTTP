package service

import (
	"task-process-service/internal/domain"
	"task-process-service/internal/repository"
)

type TaskService interface {
	AddTask(req domain.TaskAddReq) (int, error)
	UpdateTask(req domain.TaskUpdateReq) (int, error)
	DeleteTask(req domain.TaskDeleteReq) (int, error)
	GetTaskById(req domain.TaskGetByIdReq) (domain.TaskGetDataList, error)
	GetAllTask() ([]domain.TaskGetDataList, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: repo,
	}
}

func (s *taskService) AddTask(req domain.TaskAddReq) (int, error) {
	if req.CreateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("AddTask", "CreateUserId is required, your User Id is {{req.CreateUserId}}")
	}
	id, err := s.taskRepo.Create(req)
	return id, err
}

func (s *taskService) UpdateTask(req domain.TaskUpdateReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("UpdateTask", "UpdateUserId is required, your User Id is {{req.UpdateUserId}}")
	}
	return s.taskRepo.Update(req)
}

func (s *taskService) DeleteTask(req domain.TaskDeleteReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("DeleteTask", "UpdateUserId is required, your User Id is {{req.UpdateUserId}}")
	}
	return s.taskRepo.Delete(req)
}

func (s *taskService) GetTaskById(req domain.TaskGetByIdReq) (domain.TaskGetDataList, error) {
	if req.Id <= 0 {
		return domain.TaskGetDataList{}, domain.NewIdLessThanZeroError("GetTaskById", "UserId is required, your User Id is {{req.UserId}}")
	}
	return s.taskRepo.GetById(req)
}

func (s *taskService) GetAllTask() ([]domain.TaskGetDataList, error) {
	return s.taskRepo.GetAll()
}
