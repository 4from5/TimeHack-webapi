package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Get(userId int) (webApi.UsernameInfo, error) {
	fmt.Println("service.UserService.GetAll: get", userId)
	return s.repo.Get(userId)
}
