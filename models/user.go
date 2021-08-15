package models

type User struct {
	Id       int    `json:"id" validate:"required" sql:"AUTO_INCREMENT"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
