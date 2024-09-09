package usecases

import (
	"golang-rest-api/models"
	"golang-rest-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse, error)
}

type userUsecase struct {
	ur repositories.IUserRepository
}

func NewUserUsecase(ur repositories.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user models.User) (models.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.UserResponse{}, err
	}

	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.Create(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}
