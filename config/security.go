package config

import (
	"errors"
	"os"

	"go.uber.org/zap"
)

var (
	ErrInvalidBcryptHashCost                 = errors.New("bcrypt_hash_cost must be greater than 0")
	ErrSuperuserLoginNotProvided             = errors.New("super user is enabled but his login is not provided")
	ErrSuperuserPasswordNotProvided          = errors.New("super user is enabled but his password is not set")
	ErrorAttributedRoleOnRegisterDoesntExist = errors.New("attributed role on register doesnt exist in casbin policy model")
)

type SecurityConfig struct {
	BcryptHashCost int `yaml:"bcrypt_hash_cost"`
	//Superuser creds
	SuperUserIsEnabled bool   `yaml:"enable_su"`
	SuperUserLogin     string `yaml:"su_login"`
	SuperUserPass      string `yaml:"su_password"`

	//Register config
	EnableLocalRegister      bool   `yaml:"enable_local_register"`
	AttributedRoleOnRegister string `yaml:"attributed_role_on_register"`
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
	if !Configuration.Rbac.IsRoleValid(securityConfig.AttributedRoleOnRegister) {
		Logger.Fatal("invalid attributed role on register", zap.Error(ErrorAttributedRoleOnRegisterDoesntExist))
	}
	return securityConfig
}
