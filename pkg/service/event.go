package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(userId int, event webApi.Event) (int, error) {
	fmt.Println("service.EventService.Create: userId, event", userId, " ", event)
	return s.repo.Create(userId, event)
}
func (s *EventService) GetAll(userId int) ([]webApi.Event, error) {
	fmt.Println("service.EventService.GetAll: get", userId)
	return s.repo.GetAll(userId)
}

func (s *EventService) GetById(userId int, id int) (webApi.Event, error) {
	fmt.Println("service.EventService.GetById: userId, id:", userId, ' ', id)
	return s.repo.GetById(userId, id)
}
