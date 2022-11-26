package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(userId int) ([]webApi.Task, error) {
	fmt.Println("service.TaskService.GetAll: get", userId)
	return s.repo.GetAll(userId)
}

func (s *TaskService) GetById(userId, id int) (webApi.Task, error) {
	fmt.Println("service.TaskService.GetById: get", userId, id)
	return s.repo.GetById(userId, id)
}

func (s *TaskService) Create(userId int, task webApi.Task) (int, error) {
	fmt.Println("service.TaskService.Create: get", userId, " ", task)
	return s.repo.Create(userId, task)
}
func (s *TaskService) Delete(userId, id int) error {
	fmt.Println("service.TaskService.Delete: userId, id:", userId, " ", id)
	return s.repo.Delete(userId, id)
}
