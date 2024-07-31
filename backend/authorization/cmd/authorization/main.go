package main

import (
	"log"

	"github.com/Cthulhu-tech/microservice/internal/application/services"
	"github.com/Cthulhu-tech/microservice/internal/infastructure/db/postgres"
	"github.com/Cthulhu-tech/microservice/internal/interface/api/rest"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main () {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	gormDB, err := postgres.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gormDB.AutoMigrate(&postgres.User{})

	err = gormDB.Migrator().DropTable("user")
	if err != nil {
		log.Fatal("failed to drop tables: ", err)
	}

	userRepo := postgres.NewGormUserRepository(gormDB)

	userService := services.NewUserService(userRepo)

	e := echo.New()
	rest.NewUserController(e, userService)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
