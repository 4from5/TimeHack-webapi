package service

import (
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user webApi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Categories interface {
	Create(userId int, category webApi.Category) (int, error)
	GetAll(userId int) ([]webApi.Category, error)
	GetById(userId, id int) (webApi.Category, error)
}

type Service struct {
	Authorization
	Categories
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoryService(repos.Category),
	}
}
