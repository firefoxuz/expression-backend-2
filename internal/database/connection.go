package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	dbConnection *sqlx.DB
)

func GetConnection() (*sqlx.DB, error) {
	if dbConnection == nil {
		dbSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.name"), viper.GetString("database.username"), viper.GetString("database.password"))

		db, err := sqlx.Connect("postgres", dbSourceName)
		if err != nil {
			return nil, err
		}

		dbConnection = db
	}

	return dbConnection, nil
}
