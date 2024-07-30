package request

import (
	"github.com/Cthulhu-tech/microservice/internal/application/command"
	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	Id			uuid.UUID 	`json:id`
	Email 		string 		`json:"email"`
	Name 		string 		`json:"name"`
	Surname 	string 		`json:"surname"`
	Lastname 	string 		`json:"lastname"`
	Password	string 		`json:"password"`
}

func (req *UpdateUserRequest) ToUpdateUserCommand() (*command.UpdateUserCommand, error) {
	return &command.UpdateUserCommand{
		Id: 		req.Id,
		Email: 		req.Email,
		Name: 		req.Name,
		Surname: 	req.Surname,
		Lastname: 	req.Lastname,
		Password: 	req.Password,
	}, nil
}
