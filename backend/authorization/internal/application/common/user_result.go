package common

import (
	"github.com/google/uuid"
	"time"
)

type UserResult struct {
	Id 			uuid.UUID
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Email 		string
	Name 		string
	Surname 	string
	Lastname 	string
}
