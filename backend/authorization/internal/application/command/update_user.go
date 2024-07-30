package command

import (
	"github.com/Cthulhu-tech/microservice/internal/application/common"
	"github.com/google/uuid"
)

type UpdateUserCommand struct {
	Id			uuid.UUID
	Email 		string
	Name 		string
	Surname 	string
	Lastname 	string
	Password	string
}

type UpdateUserCommandResult struct {
	Result *common.UserResult
}
