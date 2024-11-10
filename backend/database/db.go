package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func ConnectDatabase() (*gorm.DB, error) {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)
	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	db.LogMode(true) // Log SQL queries
	return db, nil
}
