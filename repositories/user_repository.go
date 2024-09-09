package repositories

import (
	"golang-rest-api/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(user *models.User, email string) error
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Create(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
