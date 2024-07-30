package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
	Ssl	   string
}

func GetDBConfig() (*DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New(err.Error())
	}

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbUsername := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbSsl := os.Getenv("SSLMODE")
	
	config := &DBConfig{
		Host:       dbHost,
		Port:       dbPort,
		Username:   dbUsername,
		Password:   dbPassword,
		Name:       dbName,
		Ssl: 		dbSsl,
	}

    return config, nil
}
