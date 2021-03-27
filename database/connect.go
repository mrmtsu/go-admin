package database

import (
	"go-admin/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	database, err := gorm.Open(mysql.Open(os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
