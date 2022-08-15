package repository

import (
	"fmt"
	"sign_in/database"
	"sign_in/model"
	"strings"

	"github.com/sirupsen/logrus"
)

type ToysPostgres struct {
	db *database.DB
}

func NewToysPostgres(db *database.DB) *ToysPostgres {
	return &ToysPostgres{db: db}
}

func (r *ToysPostgres) CreateToysRepo(userId int, toys model.Toys) (int, error) {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createToysQuery := fmt.Sprintf("INSERT INTO %s (person, movie) VALUES ($1, $2) RETURNING id", toystable)
	row := tx.QueryRow(createToysQuery, toys.Person, toys.Movie)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersToysQuery := fmt.Sprintf("INSERT INTO %s (users_id, toys_id) VALUES ($1, $2)", userstoystable)
	_, err = tx.Exec(createUsersToysQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ToysPostgres) GetAllToysRepo(userId int) ([]model.Toys, error) {
	var toys []model.Toys

	query := fmt.Sprintf("SELECT toys.id, toys.person, toys.movie FROM %s toys INNER JOIN %s users_toys on toys.id = users_toys.toys_id WHERE users_toys.users_id = $1",
		toystable, userstoystable)
	err := r.db.Conn.Select(&toys, query, userId)

	return toys, err
}

func (r *ToysPostgres) GetByIdToysRepo(userId, toysId int) (model.Toys, error) {
	var toys model.Toys

	query := fmt.Sprintf(`SELECT toys.id, toys.person, toys.movie FROM %s toys
								INNER JOIN %s users_toys on toys.id = users_toys.toys_id WHERE users_toys.users_id = $1 AND users_toys.toys_id = $2`,
		toystable, userstoystable)
	err := r.db.Conn.Get(&toys, query, userId, toysId)

	return toys, err
}

func (r *ToysPostgres) DeleteToysRepo(userId, toysId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.toys_id AND ul.users_id=$1 AND ul.toys_id=$2",
		toystable, userstoystable)
	_, err := r.db.Conn.Exec(query, userId, toysId)

	return err
}

func (r *ToysPostgres) UpdateToysRepo(userId, toysId int, input model.UpdateToysInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if &input.Person != nil {
		setValues = append(setValues, fmt.Sprintf("person=$%d", argId))
		args = append(args, *&input.Person)
		argId++
	}

	if &input.Movie != nil {
		setValues = append(setValues, fmt.Sprintf("movie=$%d", argId))
		args = append(args, *&input.Movie)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s toys SET %s FROM %s users_toys WHERE toys.id = users_toys.toys_id AND users_toys.toys_id=$%d AND users_toys.users_id=$%d",
		toystable, setQuery, userstoystable, argId, argId+1)
	args = append(args, toysId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Conn.Exec(query, args...)
	return err
}
