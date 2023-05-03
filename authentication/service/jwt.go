package authentication_service

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
	"fmt"

	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func (service *AuthenticationService) GenerateJwt(userId uint, username, kind, role string, expiration time.Duration, key *rsa.PrivateKey) (tokenString string, tokenClaims *authentication_model.JwtClaims, err error) {
	tokenClaims = &authentication_model.JwtClaims{
		Kind:   kind,
		Role:   role,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtToken := jwt.NewWithClaims(config.JwtSignMethod, tokenClaims)
	tokenString, err = jwtToken.SignedString(key)
	if err != nil {
		config.Logger.Info("Failed to sign JWT", zap.Error(err))
		return "", &authentication_model.JwtClaims{}, err
	}
	return tokenString, tokenClaims, err
}

func (service *AuthenticationService) DecodeJwt(rawToken string, publicKey *rsa.PublicKey) (claims *authentication_model.JwtClaims, err error) {
	receivedToken, err := jwt.ParseWithClaims(rawToken, &authentication_model.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		config.Logger.Info("Failed to parse JWT", zap.Error(err))
		return nil, err
	}
	// Data check
	claims, ok := receivedToken.Claims.(*authentication_model.JwtClaims)
	if !ok || !receivedToken.Valid {
		config.Logger.Info("JWT token is invalid", zap.Bool("is_token_ok", ok), zap.Bool("is_token_valid", receivedToken.Valid), zap.Error(err))
		return nil, err
	}
	return claims, nil

}
