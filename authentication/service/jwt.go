package authentication_service

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func (service *AuthenticationService) GenerateJwt(userId uint, username, kind, role string, expiration time.Duration, jwtSpecs config.JwtSpecs) (tokenString string, tokenClaims *authentication_model.JwtClaims, err error) {
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
	jwtToken := jwt.NewWithClaims(jwtSpecs.SigninMethod, tokenClaims)
	switch jwtSpecs.SigninMethod {
	case config.JwtSignMethodRS512:
		tokenString, err = jwtToken.SignedString(jwtSpecs.SignKey)
	case config.JwtSignMethodHS512:
		tokenString, err = jwtToken.SignedString(jwtSpecs.SecretKeyBytes)
	}
	if err != nil {
		config.Logger.Info("Failed to sign JWT", zap.Error(err))
		return "", &authentication_model.JwtClaims{}, err
	}
	return tokenString, tokenClaims, err
}

func (service *AuthenticationService) DecodeJwt(rawToken string, jwtSpecs config.JwtSpecs) (claims *authentication_model.JwtClaims, err error) {
	var key interface{}
	switch jwtSpecs.SigninMethod {
	case config.JwtSignMethodRS512:
		key = jwtSpecs.VerifyKey
	case config.JwtSignMethodHS512:
		key = jwtSpecs.SecretKeyBytes
	default:
		config.Logger.Error("Failed to retrieve jwt key", zap.Error(fmt.Errorf("invalid algorithm in configuration :  %v", jwtSpecs.SigninMethod.Alg())))
	}
	receivedToken, err := jwt.ParseWithClaims(rawToken, &authentication_model.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		switch jwtSpecs.SigninMethod {
		case config.JwtSignMethodRS512:
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return key, nil
		case config.JwtSignMethodHS512:
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return key, nil
		}
		return nil, fmt.Errorf("invalid algorithm in configuration :  %v", jwtSpecs.SigninMethod.Alg())
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
