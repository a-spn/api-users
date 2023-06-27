package config

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

const (
	JwtContextKey   = "jwtClientData"
	RefreshDuration = time.Hour * 24 * 10 // Refresh token is valid 10 days
	AccessDuration  = time.Minute * 5     // Access token is valid 5 minutes
)

var (
	ErrInvalidJWTProtocol = errors.New("unknown jwt signin method : available methods are 'RS512' and 'HS512'")
	JwtSignMethodRS512    = jwt.SigningMethodRS512
	JwtSignMethodHS512    = jwt.SigningMethodHS512
)

type JwtConfig struct {
	RefreshToken JwtSpecs `yaml:"refresh_token"`
	AccessToken  JwtSpecs `yaml:"access_token"`
}

type JwtSpecs struct {
	// Config
	JwtSignInMethodName string `yaml:"jwt_sign_method"`
	PrivateCertPath     string `yaml:"private_cert"`
	PublicCertPath      string `yaml:"public_cert"`
	SecretKey           string `yaml:"secret_key"`
	// Data
	SigninMethod jwt.SigningMethod

	SignKey        *rsa.PrivateKey
	VerifyKey      *rsa.PublicKey
	SecretKeyBytes []byte

	Duration time.Duration
}

func (jwtConfig *JwtConfig) InitJWT() *JwtConfig {
	jwtConfig.AccessToken = *jwtConfig.AccessToken.InitTokenConfig(AccessDuration)
	jwtConfig.RefreshToken = *jwtConfig.RefreshToken.InitTokenConfig(RefreshDuration)
	return jwtConfig
}

func (jwtSpecs *JwtSpecs) InitTokenConfig(duration time.Duration) *JwtSpecs {
	jwtSpecs.Duration = duration
	switch jwtSpecs.JwtSignInMethodName {
	case "RS512":
		jwtSpecs.SigninMethod = JwtSignMethodRS512
		jwtSpecs.SignKey = LoadPrivateKey(jwtSpecs.PrivateCertPath)
		jwtSpecs.VerifyKey = LoadPublicKey(jwtSpecs.PublicCertPath)
	case "HS512":
		jwtSpecs.SigninMethod = JwtSignMethodHS512
		if jwtSpecs.SecretKey == "" {
			Logger.Fatal("Invalid secret key", zap.Error(errors.New("jwt sign key is empty")))
		}
		jwtSpecs.SecretKeyBytes = []byte(jwtSpecs.SecretKey)
	default:
		Logger.Fatal("Invalid jwt signin method", zap.Error(ErrInvalidJWTProtocol))
	}
	return jwtSpecs
}

func LoadPublicKey(path string) *rsa.PublicKey {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		Logger.Fatal("can't load public key", zap.Error(err))
	}
	public, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		Logger.Fatal("can't parse public key", zap.Error(err))
	}
	return public

}
func LoadPrivateKey(path string) *rsa.PrivateKey {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		Logger.Fatal("can't load private key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		Logger.Fatal("can't parse private key", zap.Error(err))
	}
	return privateKey
}
