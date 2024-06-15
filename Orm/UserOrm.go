package Orm

import (
	"Authentication/Configs"
	. "Authentication/Models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SaveUser(u *User) *gorm.DB {
	u.ID = uuid.New().String()
	database := Configs.GetDBInstance()
	result := database.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	return database
}

func FindUserByEmail(email string) bool {
	Configs.InitDb()
	database := Configs.GetDBInstance()
	var user User
	if err := database.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

func FindUser(email string) (*User, error) {
	database := Configs.GetDBInstance()
	u := User{}
	if err := database.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func FindAllUser() (users []User, err error) {
	Configs.InitDb()
	db := Configs.GetDBInstance()
	var user []User
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByUserId(uId string) (*User, error) {
	Configs.InitDb()
	db := Configs.GetDBInstance()
	var user User
	if err := db.Where("id = ?", uId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
