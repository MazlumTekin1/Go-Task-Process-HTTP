package service

import (
	"strconv"
	"task-process-service/internal/domain"
	"task-process-service/internal/repository"
)

type UserService interface {
	Create(req domain.UserAddReq) (int, error)
	Update(req domain.UserUpdateReq) (int, error)
	Delete(req domain.UserDeleteReq) (int, error)
	GetById(req domain.UserGetByIdReq) (domain.UserGetDataList, error)
	GetAll() ([]domain.UserGetDataList, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) Create(req domain.UserAddReq) (int, error) {
	if req.CreateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Create", `CreateUserId has to be more than zero, your User Id is `+strconv.Itoa(req.CreateUserId)+``)
	}
	return s.userRepo.Create(req)
}

func (s *userService) Update(req domain.UserUpdateReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Update", `UpdateUserId has to be more than zero, your User Id is `+strconv.Itoa(req.UpdateUserId)+``)
	}
	return s.userRepo.Update(req)
}

func (s *userService) Delete(req domain.UserDeleteReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("Delete", `UpdateUserId has to be more than zero, your User Id is `+strconv.Itoa(req.UpdateUserId)+``)
	}
	return s.userRepo.Delete(req)
}

func (s *userService) GetById(req domain.UserGetByIdReq) (domain.UserGetDataList, error) {
	if req.Id <= 0 {
		return domain.UserGetDataList{}, domain.NewIdLessThanZeroError("GetById", `UserId has to be more than zero, your User Id is `+strconv.Itoa(req.Id)+``)
	}
	return s.userRepo.GetById(req)
}

func (s *userService) GetAll() ([]domain.UserGetDataList, error) {
	return s.userRepo.GetAll()
}
