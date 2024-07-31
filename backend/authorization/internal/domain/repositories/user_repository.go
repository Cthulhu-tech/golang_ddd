package repositories

import (
	"github.com/Cthulhu-tech/microservice/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entities.ValidatedUser) (*entities.User, error)
	Update(user *entities.ValidatedUser) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByName(name string) ([]*entities.User, error)
	FindBySurname(surname string) ([]*entities.User, error)
	FindByLastname(lastname string) ([]*entities.User, error)
	Delete(id uuid.UUID) error
	FindAll(offset, count int) ([]*entities.User, error)
}
