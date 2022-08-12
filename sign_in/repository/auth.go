package repository

import (
	"fmt"
	"sign_in/database"
	"sign_in/model"


)

type AuthPostgres struct {
	db *database.DB
}

func NewAuthPostgres(db *database.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", userstable)

	row := r.db.Conn.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", userstable)
	err := r.db.Conn.Get(&user, query, username, password)

	return user, err
}
