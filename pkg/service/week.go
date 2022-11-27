package service

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/4from5/TimeHack-webapi/pkg/repository"
)

type WeekService struct {
	repo repository.Week
}

func NewWeekService(repo repository.Week) *WeekService {
	return &WeekService{repo: repo}
}

func (s *WeekService) GetDays(userId int, input webApi.WeekRequest) ([]webApi.Day, error) {
	fmt.Println("service.WeekService", userId, " ", input)
	return s.repo.GetDays(userId, input)
}
