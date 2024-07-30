package postgres

import (
	"github.com/Cthulhu-tech/microservice/internal/domain/entities"
	"github.com/Cthulhu-tech/microservice/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) Create(user *entities.ValidatedUser) (*entities.User, error) {
	dbUser := toDBUser(user)

	if err := repo.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return repo.FindById(dbUser.Id)
}

func (repo *GormUserRepository) FindById(id uuid.UUID) (*entities.User, error) {
	var dbUser User
	if err := repo.db.First(&dbUser, id).Error; err != nil {
		return nil, err
	}
	return fromDBUser(&dbUser), nil
}

func (repo *GormUserRepository) FindAll(offset, count int) ([]*entities.User, error) {
	var dbUsers []User
	if err := repo.db.Limit(count).Find(&dbUsers).Offset(offset).Error; err != nil {
		return nil, err
	}

	users := make([]*entities.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = fromDBUser(&dbUser)
	}

	return users, nil
}

func (repo *GormUserRepository) Update(user *entities.ValidatedUser) (*entities.User, error) {
	dbUser := toDBUser(user)

	err := repo.db.Model(&User{}).Where("id = ?", dbUser.Id).Updates(dbUser).Error
	if err != nil {
		return nil, err
	}

	return repo.FindById(dbUser.Id)
}

func (repo *GormUserRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&User{}, id).Error
}
