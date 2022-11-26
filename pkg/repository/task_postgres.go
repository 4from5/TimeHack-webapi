package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) GetAll(userId int) ([]webApi.Task, error) {
	var Tasks []webApi.Task
	fmt.Println("repository.TaskPostgres.GetAll: get", userId)

	getAllTasks := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", tasksTable)
	err := r.db.Select(&Tasks, getAllTasks, userId)
	return Tasks, err
}

func (r *TaskPostgres) GetById(userId int, id int) (webApi.Task, error) {
	var task webApi.Task
	fmt.Println("repository.TaskPostgres.GetAll: get", userId)

	getTaskById := fmt.Sprintf("SELECT * FROM %s WHERE task_id = $1 AND user_id = $2", tasksTable)
	err := r.db.Get(&task, getTaskById, id, userId)
	return task, err
}

func (r *TaskPostgres) Create(userId int, task webApi.Task) (int, error) {
	var returnedId int
	fmt.Println("repository.TaskPostgres.Create: get", userId, task)
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id, title, description, deadline, date_time, priority) VALUES($1,$2,$3,$4,$5,$6,$7 ) RETURNING task_id", tasksTable)
	row := r.db.QueryRow(createTaskQuery, userId, task.CategoryId, task.Title, task.Description, task.Deadline, task.DateTime, task.Priority)
	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}
	return returnedId, nil

}
