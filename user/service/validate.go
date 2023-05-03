package user_service

import (
	"api-users/config"
	user_model "api-users/user/model"
	"errors"
	"net/mail"
	"regexp"
	"unicode"

	"go.uber.org/zap"
)

var (
	ErrorPasswordContainSpace = errors.New("password must not contain white space")
	ErrorPasswordRequirements = errors.New("password doesnt match minimal security requirements")
	ErrorInvalidRole          = errors.New("given role doesnt exist")
	ErrorInvalidEmailAdress   = errors.New("email address is invalid")
	ErrorInvalidUsername      = errors.New("username is invalid")
)

func (service *UserService) ValidateUser(user user_model.User, allowEmptyFields bool) error {
	if !(allowEmptyFields && user.Username == "") {
		if res, err := regexp.MatchString(`^[a-zA-Z0-9_]{3,20}$`, user.Username); !res || err != nil || user.Username == config.Configuration.Security.SuperUserLogin {
			if err != nil {
				config.Logger.Info("username is invalid", zap.Error(err))
			}
			return ErrorInvalidUsername
		}
	}
	//Check Mail
	if !(allowEmptyFields && user.Email == "") {
		if _, err := mail.ParseAddress(user.Email); err != nil {
			config.Logger.Info("email address is invalid", zap.Error(err))
			return ErrorInvalidEmailAdress
		}
	}
	//Check roles
	if !(allowEmptyFields && user.Role == "") {
		if err := service.ValidateRole(user.Role); err != nil {
			return err
		}
	}
	//Check password
	if !(allowEmptyFields && user.Password == "") {
		if err := service.ValidatePasswordRequirements(user.Password); err != nil {
			return err
		}
	}
	return nil
}

func (service UserService) ValidatePasswordRequirements(pass string) error {
	letters := 0
	sevenOrMore, number, upper, special := false, false, false, false
	for _, c := range pass {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c):
			letters++
		case c == ' ':
			config.Logger.Info("Invalid password for user creation", zap.Error(ErrorPasswordContainSpace))
			return ErrorPasswordContainSpace
		default:
		}
	}
	sevenOrMore = letters >= 7
	if !sevenOrMore || !number || !upper || !special {
		config.Logger.Info("Invalid password for user creation", zap.Error(ErrorPasswordRequirements))
		return ErrorPasswordRequirements
	}
	return nil
}

func (service UserService) ValidateRole(role string) error {
	for _, r := range service.AuthorizationService.ListAllRoles() {
		if role == r {
			return nil
		}
	}
	config.Logger.Info("Invalid role", zap.String("role", role), zap.Error(ErrorInvalidRole))
	return ErrorInvalidRole
}
