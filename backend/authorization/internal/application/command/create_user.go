package command

import "github.com/Cthulhu-tech/microservice/internal/application/common"

type CreateUserCommand struct {
	Email 		string
	Name 		string
	Surname 	string
	Lastname 	string
	Password	string
}

type CreateUserCommandResult struct {
	Result *common.UserResult
}
