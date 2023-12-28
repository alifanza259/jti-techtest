package storage

import (
	"fmt"

	"github.com/alifanza259/jwt-techtest/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config util.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
