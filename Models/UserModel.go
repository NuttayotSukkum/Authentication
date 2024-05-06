package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type UserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
