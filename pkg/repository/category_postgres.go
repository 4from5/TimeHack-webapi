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

func (r *CategoryPostgres) GetAll(userId int) ([]webApi.Category, error) {
	var categories []webApi.Category
	fmt.Println("repository.CategoryPostgres.GetAll: get", userId)

	getAllCategories := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", categoriesTable)
	err := r.db.Select(&categories, getAllCategories, userId)
	return categories, err
}
func (r *CategoryPostgres) GetById(userId, id int) (webApi.Category, error) {
	var category webApi.Category
	fmt.Println("repository.CategoryPostgres.GetById: userId, id:", userId, " ", id)

	getCategoryById := fmt.Sprintf("SELECT * FROM %s WHERE category_id = $1 AND user_id = $2", categoriesTable)
	err := r.db.Get(&category, getCategoryById, id, userId)

	return category, err
}
