package authentication_service

import (
	"api-users/config"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (service AuthenticationService) VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		config.Logger.Info("Password verification failed", zap.Error(err))
		return false
	}
	return true
}
