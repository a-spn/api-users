package user_service

import (
	"api-users/config"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (service UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Configuration.Security.BcryptHashCost)
	if err != nil {
		config.Logger.Error("Failed to hash password", zap.Error(err))
	}
	return string(bytes), err
}
