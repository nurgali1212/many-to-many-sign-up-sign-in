package repository

import (
	"fmt"
	"sign_in/database"
	"sign_in/model"
	"strings"
)

type CategoryPostgres struct {
	db *database.DB
}

func NewCategoryPostgres(db *database.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategoryRepo(toysId int, category model.Category) (int, error) {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return 0, err
	}

	var categoryId int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (genre) values ($1) RETURNING id", categorytable)

	row := tx.QueryRow(createCategoryQuery, category.Genre)
	err = row.Scan(&categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createToysCategoryQuery := fmt.Sprintf("INSERT INTO %s (toys_id, category_id) values ($1, $2)", toyscategorytable)
	_, err = tx.Exec(createToysCategoryQuery, toysId, categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return categoryId, tx.Commit()
}

func (r *CategoryPostgres) GetAllCategoryRepo(userId, toysId int) ([]model.Category, error) {
	var category []model.Category
	query := fmt.Sprintf(`SELECT category.id, category.genre, category.done FROM %s category INNER JOIN %s toyscategory on toyscategory.category_id = category.id
									INNER JOIN %s userstoys on userstoys.toys_id = toyscategory.toys_id WHERE toyscategory.toys_id = $1 AND userstoys.user_id = $2`,
		categorytable, toyscategorytable, userstoystable)
	if err := r.db.Conn.Select(&category, query, toysId, userId); err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryPostgres) GetByIdCategoryRepo(userId, categoryId int) (model.Category, error) {
	var category model.Category
	query := fmt.Sprintf(`SELECT category.id, category.genre, category.done FROM %s category INNER JOIN %s toyscategory on toyscategory.category_id = category.id
									INNER JOIN %s userstoys on userstoys.toys_id = toyscategory.toys_id WHERE category.id = $1 AND userstoys.user_id = $2`,
		categorytable, toyscategorytable, userstoystable)
	if err := r.db.Conn.Get(&category, query, categoryId, userId); err != nil {
		return category, err
	}

	return category, nil
}

func (r *CategoryPostgres) DeleteCategoryRepo(userId, categoryId int) error {
	query := fmt.Sprintf(`DELETE FROM %s category USING %s toyscategory, %s userstoys 
									WHERE category.id = toyscategory.category_id AND toyscategory.toys_id = userstoys.toys_id AND userstoys.user_id = $1 AND toyscategory.id = $2`,
		categorytable, toyscategorytable, userstoystable)
	_, err := r.db.Conn.Exec(query, userId, categoryId)
	return err
}

func (r *CategoryPostgres) UpdateCategoryRepo(userId, categoryId int, input model.UpdateCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if &input.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *&input.Genre)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s category SET %s FROM %s toyscategory, %s userstoys
									WHERE category.id = category.category_id AND toyscategory.toys_id = userstoys.toys_id AND userstoys.user_id = $%d AND category.id = $%d`,
		categorytable, setQuery, toyscategorytable, userstoystable, argId, argId+1)
	args = append(args, userId, categoryId)

	_, err := r.db.Conn.Exec(query, args...)
	return err
}
