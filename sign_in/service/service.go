package service

import (
	"sign_in/model"
	"sign_in/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}





type ToysList interface {
	CreateToysService(userId int, toys model.Toys) (int, error)
	GetAllToysService(userId int) ([]model.Toys, error)
	GetByIdToysService(userId, toysId int) (model.Toys, error)
	DeleteToysService(userId, toysId int) error
	UpdateToysService(userId, toysId int, input model.UpdateToysInput) error
}



type CreateToysInput struct {
	Person string `json:"person" db:"person"`
	Movie  string `json:"movie" db:"movie"`
}


type CategoryList interface {
	CreateCategoryService(userId, toysId int, category model.Category) (int, error)
	GetAllCategoryService(userId, toysId int) ([]model.Category, error)
	GetByIdCategoryService(userId, categoryId int) (model.Category, error)
	DeleteCategoryService(userId, categoryId int) error
	UpdateCategoryService(userId, categoryId int, input model.UpdateCategoryInput) error
}

type UpdateCategoryInput struct {
	Genre string `json:"genre"`
}

type CreateCategoryInput struct {
	Genre string `json:"genre"`
}



type Service struct {
	Authorization
	ToysList
	CategoryList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToysList:      NewToysService(repos.ToysList),
		CategoryList: NewCategoryService(repos.CategoryList,repos.ToysList),
	}
}
