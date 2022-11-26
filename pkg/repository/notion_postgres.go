package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
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
