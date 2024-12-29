package db

import (
	"log"

	"github.com/deyweson/go-gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarDB() {
	stringConexao := "host=localhost user=root password=root dbname=root port=8081 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Aconteceu um erro ao se conectar no banco!")
	}
	DB.AutoMigrate(&models.Aluno{})
}
