package usecases

import (
	"fmt"
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
	SignUp(c *gin.Context, user models.User) error
	LogIn(c *gin.Context, user models.User) (string, error)
	ConfirmEmail(c *gin.Context, tokenString string) error
}

type userUsecase struct {
	ur repositories.IUserRepository
}

func NewUserUsecase(ur repositories.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(c *gin.Context, user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logger(c).Error(err.Error())
		return err
	}

	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.Create(&newUser); err != nil {
		logger(c).Error(err.Error())
		return err
	}

	var confirmToken string
	if config.GetConfig().EmailVerification.Enabled {
		// メール確認のトークン作成
		confirmToken, err = generateComfirmToken(c, newUser.ID)
		if err != nil || confirmToken == "" {
			logger(c).Error(err.Error())
			return err
		}
		logger(c).Debug("confirmToken: " + confirmToken)
		// メール送信（非同期処理にする）
	}

	return nil
}

func (uu *userUsecase) LogIn(c *gin.Context, user models.User) (string, error) {
	storedUser := models.User{}
	if err := uu.ur.FindByEmail(&storedUser, user.Email); err != nil {
		logger(c).Error(err.Error())
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		logger(c).Error(err.Error())
		return "", err
	}

	authToken, err := generateAuthToken(c, user.ID)
	if err != nil || authToken == "" {
		logger(c).Error(err.Error())
		return "", err
	}

	return authToken, nil
}

func (uu *userUsecase) ConfirmEmail(c *gin.Context, tokenString string) error {
	// トークンを検証してユーザーIDを取得
	userID, err := validateConfirmToken(c, tokenString)
	if err != nil || userID == nil {
		logger(c).Error(err.Error())
		return err
	}

	user := models.User{}
	if err = uu.ur.FindByID(&user, *userID); err != nil {
		return err
	}

	if user.ConfirmedAt != nil {
		return fmt.Errorf("Confirmed user")
	}

	now := time.Now()
	user.ConfirmedAt = &now
	if err = uu.ur.Update(&user); err != nil {
		return err
	}

	return nil
}

// Private method

func logger(c *gin.Context) *zap.Logger {
	loggerInterface, exists := c.Get("logger")
	if !exists {
		panic("logger not found in context")
	}

	return loggerInterface.(*zap.Logger)
}

func generateAuthToken(c *gin.Context, ID uint) (string, error) {
	exp := config.GetConfig().AuthToken.TokenExpirationHours
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

func generateComfirmToken(c *gin.Context, ID uint) (string, error) {
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

func validateConfirmToken(c *gin.Context, tokenString string) (*uint, error) {
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// JWT トークンを解析してクレームを取得
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// トークンからユーザーIDを取得
	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		if exp, ok := (*claims)["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			return nil, fmt.Errorf("Token is expired")
		}
		if id, ok := (*claims)["id"].(float64); ok {
			userID := uint(id)
			return &userID, nil
		}
	}

	return nil, fmt.Errorf("Invalid token claims")
}
