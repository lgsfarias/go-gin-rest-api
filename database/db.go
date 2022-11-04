package database

import (
	"github.com/lgsfarias/go-gin-rest-api/models"
	"github.com/lgsfarias/go-gin-rest-api/setup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := setup.GetDotEnv("DB_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.Student{})
}
