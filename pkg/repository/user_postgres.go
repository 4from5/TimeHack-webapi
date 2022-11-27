package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Get(userId int) (webApi.UsernameInfo, error) {
	var username webApi.UsernameInfo
	fmt.Println("repository.UserPostgres.Get: get", userId)

	getUsernameById := fmt.Sprintf("SELECT username FROM %s WHERE user_id = $1", usersTable)
	err := r.db.Get(&username, getUsernameById, userId)
	return username, err
}
