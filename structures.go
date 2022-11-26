package webApi

import "time"

type Category struct {
	Id     int    `json:"id" db:"category_id,omitempty"`
	UserId int    `json:"user_id" db:"user_id" binding:"required"`
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
	Location           string    `json:"location" db:"location"`
	RepeatPeriod       int       `json:"repeat_period" db:"repeat_period"`
	EndPeriodTimestamp time.Time `json:"end_period_timestamp" db:"end_period_timestamp"`
}

type Task struct {
	Id          int       `json:"id" db:"task_id,omitempty"`
	UserId      int       `json:"user_id" db:"user_id" binding:"required"`
	CategoryId  int       `json:"category_id" db:"category_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	Deadline    time.Time `json:"deadline" db:"deadline"`
	DateTime    time.Time `json:"dateTime" db:"date_time" binding:"required"`
	Priority    int       `json:"priority" db:"priority"`
}

type Notion struct {
	Id          int       `json:"id" db:"task_id,omitempty"`
	UserId      int       `json:"user_id" db:"user_id" binding:"required"`
	CategoryId  int       `json:"category_id" db:"category_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	NotionText  string    `json:"notion_text" db:"notion_text"`
	CreatedDate time.Time `json:"created_date" db:"created_date" binding:"required"`
	LastUpdate  time.Time `json:"last_update" db:"last_update"`
}
