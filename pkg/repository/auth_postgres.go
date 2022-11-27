package repository

import (
	"fmt"
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgress(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user webApi.User) (int, error) {
	fmt.Println("repository.AuthPostgres.CreateUser: get", user)
	var returnedId int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) 	RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&returnedId); err != nil {
		return 0, err
	}

	return returnedId, nil
}

func (r *AuthPostgres) GetUser(username, password string) (webApi.User, error) {
	fmt.Println("repository.AuthPostgres.GetUser: get", username, password)
	var user webApi.User
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE username = $1 AND password_hash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
