package Orm

import (
	"Authentication/Configs"
	models "Authentication/Models"
	"errors"
	"gorm.io/gorm"
)

func SaveUser(User *models.User) *gorm.DB {
	Configs.InitDb()
	database := Configs.GetDBInstance()
	database.Create(&User)
	return database
}

func FindUserByEmail(email string) (bool, error) {
	Configs.InitDb()
	database := Configs.GetDBInstance()
	var user models.User
	if err := database.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FindUser(email string) (*models.User, error) {
	database := Configs.GetDBInstance()
	u := models.User{}
	if err := database.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func FindAllUser() (users []models.User, err error) {
	Configs.InitDb()
	db := Configs.GetDBInstance()
	var user []models.User
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByUserId(uId uint) (*models.User, error) {
	Configs.InitDb()
	db := Configs.GetDBInstance()
	var user models.User
	if err := db.Where("id = ?", uId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
