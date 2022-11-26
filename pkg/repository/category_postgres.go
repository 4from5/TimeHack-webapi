package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) Create(userId int, category webApi.Category) (int, error) {
	var returnedId int
	fmt.Println("repository.CategoryPostgres.Create: get", userId, category)
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, colour) VALUES($1,$2,$3) RETURNING category_id", categoriesTable)
	row := r.db.QueryRow(createCategoryQuery, userId, category.Title, category.Colour)
	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}

	return returnedId, nil

}
