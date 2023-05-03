package user_model

import "gorm.io/gorm"

type APIUser struct {
	gorm.Model
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}
