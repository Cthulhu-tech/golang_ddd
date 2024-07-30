package postgres

import (
	"errors"
	"fmt"

	"github.com/Cthulhu-tech/microservice/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	config, err := config.GetDBConfig()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.Name, config.Password, config.Ssl)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dbUri,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
