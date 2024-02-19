package database

import (
	"log"

	"github.com/gmessias/api-go-money/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "usuario:senha@/nomeDoBanco?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		log.Panic("Unable to connect - Não foi possível conectar ao banco de dados MySQL.")
	}
	DB.AutoMigrate(&models.Cash{}, &models.Category{})
}
