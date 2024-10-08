package repositories

import (
	"fmt"

	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(Username string) (*entities.Regis, error)
	CreateUser(user *entities.Regis) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(Username string) (*entities.Regis, error) {
	var user entities.Regis
	if err := r.db.Where("username = ?", Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("err: %+v\n", err)
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *entities.Regis) error {
	return r.db.Create(user).Error
}
