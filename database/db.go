package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gmessias/api-go-money/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		log.Panicf("Unable to connect to MySQL database: %v", err)
	}

	autoMigrateErr := DB.AutoMigrate(&models.Cash{}, &models.Category{})
	if autoMigrateErr != nil {
		log.Panicf("Unable to migrate database: %v", autoMigrateErr)
		return
	}
}
