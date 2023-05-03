package user_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username,omitempty" gorm:"unique"`
	Email        string `json:"email,omitempty" gorm:"unique"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty" gorm:"-"`
	Role         string `json:"role,omitempty"`
}
