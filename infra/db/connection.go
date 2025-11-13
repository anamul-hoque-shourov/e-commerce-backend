package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(config *config.DbConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.SslMode,
	)
	return connectionString
}

func NewDbConnection(config *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(&config.DbConfig)
	dbConnection, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil
}
