package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *CategoryPostgres) Delete(userId, id int) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	fmt.Println("repository.CategoryPostgres.Delete(deleting tasks): userId, id:", userId, " ", id)

	deleteTasksQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1 AND user_id = $2", tasksTable)
	_, err = tx.Exec(deleteTasksQuery, id, userId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	fmt.Println("repository.CategoryPostgres.Delete(deleting notions): userId, id:", userId, " ", id)

	deleteNotionsQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1 AND user_id = $2", notionsTable)
	_, err = tx.Exec(deleteNotionsQuery, id, userId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	fmt.Println("repository.CategoryPostgres.Delete(deleting events): userId, id:", userId, " ", id)

	deleteEventsQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1 AND user_id = $2", eventsTable)
	_, err = tx.Exec(deleteEventsQuery, id, userId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	fmt.Println("repository.CategoryPostgres.Delete(deleting category): userId, id:", userId, " ", id)

	deleteCategoriesQuery := fmt.Sprintf("DELETE FROM %s WHERE category_id = $1 AND user_id = $2", categoriesTable)
	_, err = tx.Exec(deleteCategoriesQuery, id, userId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *CategoryPostgres) Update(userId, id int, input webApi.UpdateCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Colour != nil {
		setValues = append(setValues, fmt.Sprintf("colour=$%d", argId))
		args = append(args, *input.Colour)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = %d AND category_id = %d",
		categoriesTable, setQuery, argId, argId+1) //DOn'T WORK

	args = append(args, userId, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	fmt.Println(query)

	_, err := r.db.Exec(query, args...)
	return err
}
