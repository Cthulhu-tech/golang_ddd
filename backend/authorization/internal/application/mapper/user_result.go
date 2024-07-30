package mapper

import (
	"github.com/Cthulhu-tech/microservice/internal/application/common"
	"github.com/Cthulhu-tech/microservice/internal/domain/entities"
)

func NewUserResultFromValidatedEntity(user *entities.ValidatedUser) *common.UserResult {
	return NewUserResultFromEntity(&user.User)
}

func NewUserResultFromEntity(user *entities.User) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		Id:			user.Id,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		Email: 		user.Email,
		Name: 		user.Name,
		Surname: 	user.Surname,
		Lastname: 	user.Lastname,
	}
}
