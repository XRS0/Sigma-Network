package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewPostgresDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
	viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), viper.GetString("db.dbname"),
	os.Getenv("POSTGRES_PASSWORD"), viper.GetString("db.sslmode")))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}