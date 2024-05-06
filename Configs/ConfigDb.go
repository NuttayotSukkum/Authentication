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
var e error

func InitDb() {
	dsn := os.Getenv("MYSQL_DSN")
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&models.User{})
	log.Println("Database migrated")

}
func GetDBInstance() *gorm.DB {
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
