package repository

import (
	"sign_in/database"
	"sign_in/model"
)

const (
	userstable        = "users"
	toystable         = "toys"
	userstoystable    = "users_toys"
	categorytable     = "category"
	toyscategorytable = "toys_category"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type ToysList interface {
	CreateToysRepo(userId int, toys model.Toys) (int, error)
	GetAllToysRepo(userId int) ([]model.Toys, error)
	GetByIdToysRepo(userId, toysId int) (model.Toys, error)
	DeleteToysRepo(userId, toysId int) error
	UpdateToysRepo(userId, toysId int, input model.UpdateToysInput) error
}

type CreateToysInput struct {
	Person string `json:"person" db:"person"`
	Movie  string `json:"movie" db:"movie"`
}

type CategoryList interface {
	CreateCategoryRepo(listId int, category model.Category) (int, error)
	GetAllCategoryRepo(userId, toysId int) ([]model.Category, error)
	GetByIdCategoryRepo(userId, categoryId int) (model.Category, error)
	DeleteCategoryRepo(userId, categoryId int) error
	UpdateCategoryRepo(userId, categoryId int, input model.UpdateCategoryInput) error
}

type CreateCategoryInput struct {
	Genre string `json:"genre"`
}

type Repository struct {
	Authorization
	ToysList
	CategoryList
}

type UpdateToysInput struct {
	Person string `json:"person" db:"person"`
	Movie  string `json:"movie" db:"movie"`
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToysList:      NewToysPostgres(db),
		CategoryList:  NewCategoryPostgres(db),
	}
}
