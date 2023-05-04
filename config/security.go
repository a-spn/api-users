package config

import (
	"errors"
	"os"

	"go.uber.org/zap"
)

var (
	ErrInvalidBcryptHashCost        = errors.New("bcrypt_hash_cost must be greater than 0")
	ErrSuperuserLoginNotProvided    = errors.New("super user is enabled but his login is not provided")
	ErrSuperuserPasswordNotProvided = errors.New("super user is enabled but his password is not set")
)

type SecurityConfig struct {
	BcryptHashCost int `yaml:"bcrypt_hash_cost"`
	//Superuser creds
	SuperUserIsEnabled bool   `yaml:"enable_su"`
	SuperUserLogin     string `yaml:"su_login"`
	SuperUserPass      string `yaml:"su_password"`
}

func (securityConfig *SecurityConfig) InitSecurity() *SecurityConfig {
	if securityConfig.BcryptHashCost == 0 {
		Logger.Fatal("error in configuration.", zap.Error(ErrInvalidBcryptHashCost))
	} else if securityConfig.BcryptHashCost < 12 {
		Logger.Warn("Unsafe configuration : bcrypt_hash_cost should be greater 11 for safety reasons")
	}
	if securityConfig.SuperUserIsEnabled {
		if securityConfig.SuperUserLogin == "" {
			Logger.Fatal("error in configuration", zap.Error(ErrSuperuserLoginNotProvided))
		}
		if securityConfig.SuperUserPass == "" {
			if os.Getenv("SU_PASSWORD") == "" {
				Logger.Fatal("error in configuration", zap.Error(ErrSuperuserPasswordNotProvided))
			} else {
				securityConfig.SuperUserPass = os.Getenv("SU_PASSWORD")
			}
		}
	}
	return securityConfig
}
