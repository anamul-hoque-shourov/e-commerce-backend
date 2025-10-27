package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(config *config.Config) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		config.DBConfig.User,
		config.DBConfig.Password,
		config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.Name,
		config.DBConfig.SslMode,
	)
	return connectionString
}

func NewConnection(config *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(config)
	dbConnection, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil
}
