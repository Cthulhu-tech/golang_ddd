package repositories

import (
	"github.com/Cthulhu-tech/microservice/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entities.ValidatedUser) (*entities.User, error)
	Update(user *entities.ValidatedUser) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	Delete(id uuid.UUID) error
	FindAll(offset, count int) ([]*entities.User, error)
}
