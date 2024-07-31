package entities

import (
	"errors"
	"time"

	"github.com/Cthulhu-tech/microservice/internal/application/command"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	Email 		string  	`validate:"required,email"`
	Name 		string		`validate:"required,min=1"`
	Surname 	string		`validate:"omitempty,min=1"`
	Lastname 	string		`validate:"omitempty,min=1"`
	Password	string  	`validate:"required,min=8,containsany=!@#?*"`
}

func (u *User) validate() error {
	v := validator.New();
	err := v.Struct(u)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func NewUser(user *command.CreateUserCommand) *User {
	return &User{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user.Name,
		Surname:   user.Surname,
		Lastname:  user.Lastname,
		Password:  user.Password,
		Email:     user.Email,
	}
}

func (u *User) UpdateEmail(email string) error {
	u.Email = email
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) UpdateName(name string) error {
	u.Name = name
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) UpdateSurName(surname string) error {
	u.Surname = surname
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) UpdateLastName(lastname string) error {
	u.Lastname = lastname
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	u.UpdatedAt = time.Now()

	return nil
}

