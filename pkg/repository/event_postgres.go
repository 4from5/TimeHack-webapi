package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) Create(userId int, event webApi.Event) (int, error) {
	var returnedId int
	fmt.Println("repository.EventPostgres.Create: user id, event", userId, " ", event)
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id, title, description, start_timestamp,end_timestamp,is_full_day,event_location,repeat_period,end_period_timestamp) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING category_id", eventsTable)

	row := r.db.QueryRow(createCategoryQuery, userId, event.CategoryId, event.Title,
		event.Description, event.StartTimestamp, event.EndTimestamp, event.IsFullDay,
		event.EventLocation, event.RepeatPeriod, event.EndPeriodTimestamp)
	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}

	return returnedId, nil

}
func (r *EventPostgres) GetAll(userId int) ([]webApi.Event, error) {
	var events []webApi.Event
	fmt.Println("repository.EventPostgres.GetAll: get", userId)

	getAllEvents := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", eventsTable)
	err := r.db.Select(&events, getAllEvents, userId)
	return events, err
}
func (r *EventPostgres) GetById(userId int, id int) (webApi.Event, error) {
	var event webApi.Event
	fmt.Println("repository.EventPostgres.GetById: userId, id:", userId, " ", id)

	getEventById := fmt.Sprintf("SELECT * FROM %s WHERE event_id = $1 AND user_id = $2", eventsTable)
	err := r.db.Get(&event, getEventById, id, userId)

	return event, err
}
