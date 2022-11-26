package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type NotionService struct {
	repo repository.Notion
}

func NewNotionService(repo repository.Notion) *NotionService {
	return &NotionService{repo: repo}
}

func (s *NotionService) GetAll(userId int) ([]webApi.Notion, error) {
	fmt.Println("service.NotionService.GetAll: get", userId)
	return s.repo.GetAll(userId)
}

func (s *NotionService) GetById(userId, id int) (webApi.Notion, error) {
	fmt.Println("service.NotionService.GetById: get", userId, id)
	return s.repo.GetById(userId, id)
}

func (s *NotionService) Create(userId int, notion webApi.Notion) (int, error) {
	fmt.Println("service.NotionService.Create: get", userId, " ", notion)
	return s.repo.Create(userId, notion)
}
