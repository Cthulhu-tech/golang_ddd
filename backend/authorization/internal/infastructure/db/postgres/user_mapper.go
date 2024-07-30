package postgres

import "github.com/Cthulhu-tech/microservice/internal/domain/entities"


func toDBUser(user *entities.ValidatedUser) *User {
	s := &User{
		Name: user.Name,
		Surname: user.Surname,
		Lastname: user.Lastname,
		Email: user.Email,
		Id: user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Password: user.Password,
	}

	return s
}


func fromDBUser(user *User) *entities.User {
	s := &entities.User{
		Name: user.Name,
		Surname: user.Surname,
		Lastname: user.Lastname,
		Email: user.Email,
		Id: user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return s
}
