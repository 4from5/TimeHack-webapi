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

func (s *EventService) GetSchedule(userId int, group webApi.Group) ([]webApi.Event, error) {
	fmt.Println("service.EventService.Schedule: group:", group)
	events := Scraper(group)
	for i, _ := range events {
		_, err := s.repo.Create(userId, events[i])
		if err != nil {
			return nil, err
		}
	}
	return s.repo.GetAll(userId)
}

func (s *EventService) Download(userId int) {
	fmt.Println("service.EventService.Download: id:", userId)
	events, _ := s.repo.GetAll(userId)
	fmt.Println(events)
	Serialize(events)
}

func (s *EventService) Delete(userId, id int) error {
	fmt.Println("service.EventService.Delete: userId, id:", userId, " ", id)
	return s.repo.Delete(userId, id)
}

func (s *EventService) Update(userId, id int, input webApi.UpdateEventInput) error {
	fmt.Println("service.CategoryService.Update: userId, id, input:", userId, " ", id, " ", input)
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
