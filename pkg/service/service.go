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

type User interface {
	Get(userId int) (webApi.UsernameInfo, error)
}

type Categories interface {
	Create(userId int, category webApi.Category) (int, error)
	GetAll(userId int) ([]webApi.Category, error)
	GetById(userId, id int) (webApi.Category, error)
	Delete(userId, id int) error
	Update(userId, id int, input webApi.UpdateCategoryInput) error
}

type Events interface {
	Create(userId int, event webApi.Event) (int, error)
	GetAll(userId int) ([]webApi.Event, error)
	GetById(userId, id int) (webApi.Event, error)
	GetSchedule(userId int, group webApi.Group) ([]webApi.Event, error)
	Delete(userId, id int) error
	Download(userId int)
	Update(userId, id int, input webApi.UpdateEventInput) error
}
type Notions interface {
	GetAll(userId int) ([]webApi.Notion, error)
	GetById(userId, id int) (webApi.Notion, error)
	Create(userId int, notion webApi.Notion) (int, error)
	Delete(userId, id int) error
	Update(userId, id int, input webApi.UpdateNotionInput) error
}

type Tasks interface {
	GetAll(userId int) ([]webApi.Task, error)
	GetById(userId, id int) (webApi.Task, error)
	Create(userId int, task webApi.Task) (int, error)
	Delete(userId, id int) error
	Update(userId, id int, input webApi.UpdateTaskInput) error
}

type Service struct {
	Authorization
	Categories
	Events
	Notions
	Tasks
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoryService(repos.Category),
		Events:        NewEventService(repos.Event),
		Notions:       NewNotionService(repos.Notion),
		Tasks:         NewTaskService(repos.Task),
		User:          NewUserService(repos.User),
	}
}
