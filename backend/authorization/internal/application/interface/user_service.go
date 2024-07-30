package interfaces

import (
	"github.com/Cthulhu-tech/microservice/internal/application/command"
	"github.com/Cthulhu-tech/microservice/internal/application/query"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(updateCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	FindAllUsers(offset, count int) (*query.UserQueryListResult, error)
	FindUserById(id uuid.UUID) (*query.UserQueryResult, error)
	UpdateUser(updateCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error)
	DeleteUser(id uuid.UUID) error
}
