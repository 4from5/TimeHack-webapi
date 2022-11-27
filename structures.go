package webApi

import (
	"errors"
	"time"
)

type Category struct {
	Id     int    `json:"id" db:"category_id,omitempty"`
	UserId int    `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title" binding:"required"`
	Colour string `json:"colour" db:"colour"`
}

type Event struct {
	Id                 int       `json:"id" db:"event_id,omitempty"`
	UserId             int       `json:"user_id" db:"user_id" binding:"required"`
	CategoryId         int       `json:"category_id" db:"category_id" binding:"required"`
	Title              string    `json:"title" db:"title" binding:"required"`
	Description        string    `json:"description" db:"description"`
	StartTimestamp     time.Time `json:"start_timestamp" db:"start_timestamp" binding:"required"`
	EndTimestamp       time.Time `json:"end_timestamp" db:"end_timestamp" binding:"required"`
	IsFullDay          bool      `json:"is_full_day" db:"is_full_day"`
	EventLocation      string    `json:"event_location" db:"event_location"`
	RepeatPeriodDays   int       `json:"repeat_period_days" db:"repeat_period_days"`
	EndPeriodTimestamp time.Time `json:"end_period_timestamp" db:"end_period_timestamp"`
}

type Task struct {
	Id           int       `json:"id" db:"task_id,omitempty"`
	UserId       int       `json:"user_id" db:"user_id" binding:"required"`
	CategoryId   int       `json:"category_id" db:"category_id" binding:"required"`
	Title        string    `json:"title" db:"title" binding:"required"`
	Description  string    `json:"description" db:"description"`
	Deadline     time.Time `json:"deadline" db:"deadline"`
	DateTime     time.Time `json:"date_time" db:"date_time" binding:"required"`
	CreationDate time.Time `json:"creation_date" db:"creation_date" binding:"required"`
	Priority     int       `json:"priority" db:"priority"`
	IsDone       bool      `json:"is_done" db:"is_done"`
}

type Notion struct {
	Id          int       `json:"id" db:"notion_id,omitempty"`
	UserId      int       `json:"user_id" db:"user_id" binding:"required"`
	CategoryId  int       `json:"category_id" db:"category_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	NotionText  string    `json:"notion_text" db:"notion_text"`
	CreatedDate time.Time `json:"created_date" db:"created_date" binding:"required"`
	LastUpdate  time.Time `json:"last_update" db:"last_update"`
}

type Group struct {
	GroupName  string `json:"group_name"`
	CategoryId int    `json:"category_id"`
}

type UpdateCategoryInput struct {
	Title  *string `json:"title"`
	Colour *string `json:"colour"`
}

type UpdateEventInput struct {
	CategoryId         *int       `json:"category_id"`
	Title              *string    `json:"title"`
	Description        *string    `json:"description"`
	StartTimestamp     *time.Time `json:"start_timestamp"`
	EndTimestamp       *time.Time `json:"end_timestamp"`
	IsFullDay          *bool      `json:"is_full_day"`
	EventLocation      *string    `json:"event_location"`
	RepeatPeriodDays   *int       `json:"repeat_period_days"`
	EndPeriodTimestamp *time.Time `json:"end_period_timestamp"`
}

type UpdateTaskInput struct {
	CategoryId  *int       `json:"category_id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	DateTime    *time.Time `json:"date_time"`
	Priority    *int       `json:"priority"`
	IsDone      *bool      `json:"is_done"`
}

type UpdateNotionInput struct {
	CategoryId *int       `json:"category_id"`
	Title      *string    `json:"title"`
	NotionText *string    `json:"notion_text"`
	LastUpdate *time.Time `json:"last_update"`
}

type UsernameInfo struct {
	Username string `json:"username"`
}

func (i UpdateCategoryInput) Validate() error {
	if i.Title == nil && i.Colour == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateEventInput) Validate() error {
	if i.Title == nil && i.CategoryId == nil && i.Description == nil && i.EndPeriodTimestamp == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateNotionInput) Validate() error {
	if i.Title == nil && i.CategoryId == nil && i.NotionText == nil && i.LastUpdate == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateTaskInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Deadline == nil && i.CategoryId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
