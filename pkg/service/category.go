package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(userId int, category webApi.Category) (int, error) {
	fmt.Println("service.CategoryService.Create: get", userId, " ", category)
	return s.repo.Create(userId, category)
}
