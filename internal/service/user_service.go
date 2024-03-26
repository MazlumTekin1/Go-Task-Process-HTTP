package service

import (
	"task-process-service/internal/domain"
	"task-process-service/internal/repository"
)

type UserService interface {
	AddUser(req domain.UserAddReq) (int, error)
	UpdateUser(req domain.UserUpdateReq) (int, error)
	DeleteUser(req domain.UserDeleteReq) (int, error)
	GetUserById(req domain.UserGetByIdReq) (domain.UserGetDataList, error)
	GetAllUser() ([]domain.UserGetDataList, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) AddUser(req domain.UserAddReq) (int, error) {
	if req.CreateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("AddUser", "CreateUserId is required, your User Id is {{req.CreateUserId}}")
	}
	return s.userRepo.Create(req)
}

func (s *userService) UpdateUser(req domain.UserUpdateReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("UpdateUser", "UpdateUserId is required, your User Id is {{req.UpdateUserId}}")
	}
	return s.userRepo.Update(req)
}

func (s *userService) DeleteUser(req domain.UserDeleteReq) (int, error) {
	if req.UpdateUserId <= 0 {
		return 0, domain.NewUserIdIsRequiredError("DeleteUser", "UpdateUserId is required, your User Id is {{req.UpdateUserId}}")
	}
	return s.userRepo.Delete(req)
}

func (s *userService) GetUserById(req domain.UserGetByIdReq) (domain.UserGetDataList, error) {
	if req.Id <= 0 {
		return domain.UserGetDataList{}, domain.NewIdLessThanZeroError("GetUserById", "UserId is required, your User Id is {{req.UserId}}")
	}
	return s.userRepo.GetById(req)
}

func (s *userService) GetAllUser() ([]domain.UserGetDataList, error) {
	return s.userRepo.GetAll()
}
