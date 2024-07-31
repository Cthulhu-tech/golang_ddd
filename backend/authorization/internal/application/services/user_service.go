package services

import (
	"errors"
	"strings"

	"github.com/Cthulhu-tech/microservice/internal/application/command"
	interfaces "github.com/Cthulhu-tech/microservice/internal/application/interface"
	"github.com/Cthulhu-tech/microservice/internal/application/mapper"
	"github.com/Cthulhu-tech/microservice/internal/application/query"
	"github.com/Cthulhu-tech/microservice/internal/domain/entities"
	"github.com/Cthulhu-tech/microservice/internal/domain/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) interfaces.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	var newUser = entities.NewUser(user)

	validatedUser, err := entities.NewValidatedUser(newUser)
	if err != nil {
		return nil, err
	}

	findUser, err := s.repo.FindByEmail(validatedUser.Email)
	if err != nil {
		return nil, err
	}
	
	if findUser != nil {
		return nil, errors.New("user email found: " + findUser.Email)
	}

	_, err = s.repo.Create(validatedUser)
	if err != nil {
		return nil, err
	}

	result := command.CreateUserCommandResult{
		Result: mapper.NewUserResultFromValidatedEntity(validatedUser),
	}

	return &result, nil
}

func (s *UserService) FindAllUsers(offset, count int) (*query.UserQueryListResult, error) {
	storedUsers, err := s.repo.FindAll(offset, count)
	if err != nil {
		return nil, err
	}

	var queryResult query.UserQueryListResult
	for _, user := range storedUsers {
		queryResult.Result = append(queryResult.Result, mapper.NewUserResultFromEntity(user))
	}

	return &queryResult, nil
}

func (s *UserService) FindUserById(id uuid.UUID) (*query.UserQueryResult, error) {
	storedUser, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	var queryResult query.UserQueryResult
	queryResult.Result = mapper.NewUserResultFromEntity(storedUser)

	return &queryResult, nil
}

func (s *UserService) UpdateUser(updateCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	user, err := s.repo.FindById(updateCommand.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	var updateErrors []string

	if updateCommand.Name != "" {
		if err := user.UpdateName(updateCommand.Name); err != nil {
			updateErrors = append(updateErrors, err.Error())
		}
	}

	if updateCommand.Lastname != "" {
		if err := user.UpdateLastName(updateCommand.Lastname); err != nil {
			updateErrors = append(updateErrors, err.Error())
		}
	}

	if updateCommand.Surname != "" {
		if err := user.UpdateSurName(updateCommand.Surname); err != nil {
			updateErrors = append(updateErrors, err.Error())
		}
	}

	if updateCommand.Email != "" {
		if err := user.UpdateEmail(updateCommand.Email); err != nil {
			updateErrors = append(updateErrors, err.Error())
		}
	}

	if updateCommand.Password != "" {
		if err := user.UpdatePassword(updateCommand.Password); err != nil {
			updateErrors = append(updateErrors, err.Error())
		}
	}

	if len(updateErrors) > 0 {
		return nil, errors.New("Update error: " + strings.Join(updateErrors, ", "))
	}

	validatedUpdatedUser, err := entities.NewValidatedUser(user)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.Update(validatedUpdatedUser)
	if err != nil {
		return nil, err
	}

	result := command.UpdateUserCommandResult{
		Result: mapper.NewUserResultFromEntity(user),
	}

	return &result, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}