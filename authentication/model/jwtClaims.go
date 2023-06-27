package authentication_model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Kind   string `json:"kind"`
	Role   string `json:"role"`
	UserID uint   `json:"user_id"`
	jwt.RegisteredClaims
}
