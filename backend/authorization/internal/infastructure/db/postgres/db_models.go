package postgres

import (
	"github.com/google/uuid"
	"time"
)

type JWT struct {
	Id		uuid.UUID	`gorm:"primaryKey"`
	User	User		`gorm:"foreignKey:UserID"`
	Token	string
}

type User struct {
	Id 			uuid.UUID 	`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Email 		string
	Name 		string
	Surname 	string
	Lastname 	string
	Password	string
}
