package Configs

import (
	"Authentication/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var database *gorm.DB

func InitDb() {
	InitEnv()
	dsn := os.Getenv("MYSQL_DSN")
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	database = db

	log.Println("Database migrated")
}

func GetDBInstance() *gorm.DB {
	if database == nil {
		log.Fatal("Database is not initialized")
	}
	log.Println("GetDBInstance")
	return database
}

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".ENV is Start")
}
