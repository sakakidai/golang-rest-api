package usecases

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/repositories"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(c *gin.Context, user models.User) (string, error)
}

type userUsecase struct {
	ur repositories.IUserRepository
}

func NewUserUsecase(ur repositories.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(c *gin.Context, user models.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logger(c).Error(err.Error())
		return "", err
	}

	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.Create(&newUser); err != nil {
		logger(c).Error(err.Error())
		return "", err
	}

	var confirmToken string
	if config.GetConfig().EmailVerification.Enabled {
		// メール確認のトークン作成
		confirmToken, err = uu.generateComfirmToken(c, newUser.ID)
		if err != nil {
			logger(c).Error(err.Error())
			return "", err
		}
		logger(c).Debug("confirmToken: " + confirmToken)
		// メール送信（非同期処理にする）
	}

	return confirmToken, nil
}

// Private method

func logger(c *gin.Context) *zap.Logger {
	loggerInterface, exists := c.Get("logger")
	if !exists {
		panic("logger not found in context")
	}

	return loggerInterface.(*zap.Logger)
}

func (uu *userUsecase) generateComfirmToken(c *gin.Context, ID uint) (string, error) {
	exp := config.GetConfig().EmailVerification.TokenExpirationHours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(time.Hour * exp).Unix(),
	})

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		logger(c).Error(err.Error())
		return "", err
	}

	return tokenString, nil
}
