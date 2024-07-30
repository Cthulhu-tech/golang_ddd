package request

import "github.com/Cthulhu-tech/microservice/internal/application/command"


type CreateUserRequest struct {
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Surname 	string `json:"surname"`
	Lastname 	string `json:"lastname"`
	Password	string `json:"password"`
}

func (req *CreateUserRequest) ToCreateUserCommand() (*command.CreateUserCommand, error) {
	return &command.CreateUserCommand{
		Email: 		req.Email,
		Name: 		req.Name,
		Surname: 	req.Surname,
		Lastname: 	req.Lastname,
		Password:	req.Password,
	}, nil
}