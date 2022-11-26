package repository

import (
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user webApi.User) (int, error)
	GetUser(username, password string) (webApi.User, error)
}

type Repository struct {
	Authorization
	Category
	//Task
	Event
	Notion
	
}

type Category interface {
	Create(userId int, category webApi.Category) (int, error)
	GetAll(userId int) ([]webApi.Category, error)
	GetById(userId, id int) (webApi.Category, error)
}


type Event interface {
	Create(userId int, category webApi.Event) (int, error)
	GetAll(userId int) ([]webApi.Event, error)
	GetById(userId, id int) (webApi.Event, error)
}
type Notion interface {
	GetAll(userId int) ([]webApi.Notion, error)
	GetById(userId int, id int) (webApi.Notion, error)
	Create(userId int, notion webApi.Notion) (int, error)

}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgress(db),
		Category:      NewCategoryPostgres(db),
		Event:         NewEventPostgres(db),
		Notion:        NewNotionPostgres(db),
	}
}
