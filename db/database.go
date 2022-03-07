package db

import (
	"log"

	"github.com/tjaime/go-gin-api-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}

	DB.AutoMigrate(&model.Aluno{})
}
