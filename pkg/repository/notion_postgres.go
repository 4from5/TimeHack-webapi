package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type NotionPostgres struct {
	db *sqlx.DB
}

func NewNotionPostgres(db *sqlx.DB) *NotionPostgres {
	return &NotionPostgres{db: db}
}

func (r *NotionPostgres) GetAll(userId int) ([]webApi.Notion, error) {
	var notions []webApi.Notion
	fmt.Println("repository.NotionPostgres.GetAll: get", userId)

	getAllCategories := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", notionsTable)
	err := r.db.Select(&notions, getAllCategories, userId)
	return notions, err
}

func (r *NotionPostgres) GetById(userId int, id int) (webApi.Notion, error) {
	var notion webApi.Notion
	fmt.Println("repository.NotionPostgres.GetAll: get", userId)

	getNotionById := fmt.Sprintf("SELECT * FROM %s WHERE notion_id = $1 AND user_id = $2", notionsTable)
	err := r.db.Get(&notion, getNotionById, id, userId)
	return notion, err
}

func (r *NotionPostgres) Create(userId int, notion webApi.Notion) (int, error) {
	var returnedId int
	fmt.Println("repository.NotionPostgres.Create: get", userId, notion)
	createNotionQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id, title, notion_text, created_date, last_update) VALUES($1,$2,$3,$4,$5,$6) RETURNING notion_id", notionsTable)
	row := r.db.QueryRow(createNotionQuery, userId, notion.CategoryId, notion.Title, notion.NotionText, notion.CreatedDate, notion.LastUpdate)
	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}
	return returnedId, nil

}

func (r *NotionPostgres) Delete(userId int, id int) error {

	fmt.Println("repository.NotionPostgres.Delete: userId, id:", userId, " ", id)

	query := fmt.Sprintf("DELETE FROM %s WHERE event_id = $1 AND user_id = $2", eventsTable)

	_, err := r.db.Exec(query, id, userId)

	return err
}

func (r *NotionPostgres) Update(userId, id int, input webApi.UpdateNotionInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.NotionText != nil {
		setValues = append(setValues, fmt.Sprintf("notion_text=$%d", argId))
		args = append(args, *input.NotionText)
		argId++
	}

	if input.CategoryId != nil {
		setValues = append(setValues, fmt.Sprintf("category_id=$%d", argId))
		args = append(args, *input.CategoryId)
		argId++
	}

	if input.LastUpdate != nil {
		setValues = append(setValues, fmt.Sprintf("last_update=$%d", argId))
		args = append(args, *input.LastUpdate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND notion_id = $%d",
		notionsTable, setQuery, argId, argId+1)

	args = append(args, userId, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	fmt.Println(query)

	_, err := r.db.Exec(query, args...)
	return err
}
