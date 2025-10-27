package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SslMode  string
}

type Config struct {
	Version     string
	ServiceName string
	Port        int
	JwtSecret   string
	DBConfig    DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		fmt.Println("Port is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Port must be a number", err)
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT Secret is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB Port is required")
		os.Exit(1)
	}

	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("DB Port must be a number", err)
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	dbSslMode := os.Getenv("DB_SSLMODE")
	if dbSslMode == "" {
		fmt.Println("DB SslMode is required")
		os.Exit(1)
	}

	dbConfig := DBConfig{
		Host:     dbHost,
		Port:     dbPortInt,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		SslMode:  dbSslMode,
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		Port:        port,
		JwtSecret:   jwtSecret,
		DBConfig:    dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
