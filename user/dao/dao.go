package user_dao

import (
	"errors"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

var (
	ErrUserDoesNotExist  = errors.New("ressource doesnt exist")
	ErrDuplicateKeyEntry = errors.New("duplicate key entry")
)

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}
