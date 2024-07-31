package request

import (
	"errors"

	"github.com/Cthulhu-tech/microservice/internal/application/command"
	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	Id			uuid.UUID 	`json:"id"`
	Email 		string 		`json:"email,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	Surname 	string 		`json:"surname,omitempty"`
	Lastname 	string 		`json:"lastname,omitempty"`
	Password	string 		`json:"password,omitempty"`
}

func (req *UpdateUserRequest) ToUpdateUserCommand() (*command.UpdateUserCommand, error) {
	if req.Id == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}
	
	return &command.UpdateUserCommand{
		Id: 		req.Id,
		Email: 		req.Email,
		Name: 		req.Name,
		Surname: 	req.Surname,
		Lastname: 	req.Lastname,
		Password: 	req.Password,
	}, nil
}
