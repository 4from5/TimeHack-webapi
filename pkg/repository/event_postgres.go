package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id, title, description, start_timestamp,end_timestamp,is_full_day,event_location,repeat_period_days,end_period_timestamp) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING category_id", eventsTable)

	row := r.db.QueryRow(createCategoryQuery, userId, event.CategoryId, event.Title,
		event.Description, event.StartTimestamp, event.EndTimestamp, event.IsFullDay,
		event.EventLocation, event.RepeatPeriodDays, event.EndPeriodTimestamp)
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

func (r *EventPostgres) Delete(userId int, id int) error {

	fmt.Println("repository.EventPostgres.Delete: userId, id:", userId, " ", id)

	query := fmt.Sprintf("DELETE FROM %s WHERE event_id = $1 AND user_id = $2", eventsTable)

	_, err := r.db.Exec(query, id, userId)

	return err
}

func (r *EventPostgres) Update(userId, id int, input webApi.UpdateEventInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.StartTimestamp != nil {
		setValues = append(setValues, fmt.Sprintf("start_timestamp=$%d", argId))
		args = append(args, *input.StartTimestamp)
		argId++
	}

	if input.EndTimestamp != nil {
		setValues = append(setValues, fmt.Sprintf("end_timestamp=$%d", argId))
		args = append(args, *input.EndTimestamp)
		argId++
	}

	if input.IsFullDay != nil {
		setValues = append(setValues, fmt.Sprintf("is_full_day=$%d", argId))
		args = append(args, *input.IsFullDay)
		argId++
	}

	if input.EventLocation != nil {
		setValues = append(setValues, fmt.Sprintf("event_location=$%d", argId))
		args = append(args, *input.EventLocation)
		argId++
	}

	if input.RepeatPeriodDays != nil {
		setValues = append(setValues, fmt.Sprintf("repeat_period_days=$%d", argId))
		args = append(args, *input.RepeatPeriodDays)
		argId++
	}

	if input.EndPeriodTimestamp != nil {
		setValues = append(setValues, fmt.Sprintf("end_period_timestamp=$%d", argId))
		args = append(args, *input.EndPeriodTimestamp)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND event_id = $%d",
		eventsTable, setQuery, argId, argId+1)

	args = append(args, userId, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	fmt.Println(query)

	_, err := r.db.Exec(query, args...)
	return err
}
