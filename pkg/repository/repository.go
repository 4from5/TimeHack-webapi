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
	//Notion
	//Event
}

type Category interface {
	Create(userId int, category webApi.Category) (int, error)
	GetAll(userId int) ([]webApi.Category, error)
	GetById(userId, id int) (webApi.Category, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgress(db),
		Category:      NewCategoryPostgres(db),
	}
}
