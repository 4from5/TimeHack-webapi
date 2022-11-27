package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
	"time"
)

type WeekPostgres struct {
	db *sqlx.DB
}

func NewWeekPostgres(db *sqlx.DB) *WeekPostgres {
	return &WeekPostgres{db: db}
}

func (r *WeekPostgres) GetDays(userId int, input webApi.WeekRequest) ([]webApi.Day, error) {
	Week := make([]webApi.Day, 7, 7)
	fmt.Println("repository.WeekPostgres", userId)
	var i int
	for i = 0; i < 7; i++ {

		tasksQuery := fmt.Sprintf("SELECT COUNT(task_id) FROM %s WHERE (date_time = date($1)) AND not is_done AND user_id = $2", tasksTable)
		row := r.db.QueryRow(tasksQuery, input.FirstDay, userId)

		if err := row.Scan(&Week[i].TasksCount); err != nil {
			return nil, err
		}

		eventsQuery := fmt.Sprintf("SELECT COUNT(event_id) FROM %s WHERE (start_timestamp BETWEEN date($1) AND date ($2)) AND user_id = $3", eventsTable)
		row = r.db.QueryRow(eventsQuery, input.FirstDay, input.FirstDay.Add(time.Hour*24), userId)

		if err := row.Scan(&Week[i].EventsCount); err != nil {
			return nil, err
		}

		deadlinesQuery := fmt.Sprintf("SELECT COUNT(task_id) FROM %s WHERE not is_done AND user_id = $1 AND deadline = date($2)", tasksTable)
		row = r.db.QueryRow(deadlinesQuery, userId, input.FirstDay)
		if err := row.Scan(&Week[i].DeadlinesCount); err != nil {
			return nil, err
		}

		input.FirstDay = input.FirstDay.Add(time.Hour * 24)

	}
	return Week, nil
}
