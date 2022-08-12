package model

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Toys struct {
	Id     uint   `json:"id" db:"id"`
	Person string `json:"person" db:"person"`
	Movie  string `json:"movie" db:"movie"`
	// UserId int 	`json:"user_id" db:"user_id"`
}

type Category struct {
	Id    uint   `json:"id"`
	Genre string `json:"genre"`
}

type UserToys struct {
	Id     uint `json:"id" db:"id"`
	UserId uint `json:"user_id" db:"user_id"`
	ToysId uint `json:"toys_id_id" db:"toys_id"`
}

type ToysCategory struct {
	Id         uint `json:"id" db:"id"`
	ToysId     uint `json:"toys_id" db:"toys_id"`
	CategoryId uint `json:"category_id" db:"category_id"`
}

type UpdateToysInput struct {
	Person string `json:"person" db:"person"`
	Movie  string `json:"movie" db:"movie"`
}

type UpdateCategoryInput struct {
	Genre string `json:"genre"`
}

func (i *UpdateToysInput) Validate() error {
	if &i.Person == nil && &i.Movie == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
