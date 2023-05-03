package config

import (
	"crypto/rsa"
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
	JwtSignMethod = jwt.SigningMethodRS512
)

type JwtConfig struct {
	RefreshToken JwtSpecs `yaml:"refresh_token"`
	AccessToken  JwtSpecs `yaml:"access_token"`
}

type JwtSpecs struct {
	Duration        time.Duration
	PrivateCertPath string `yaml:"private_cert"`
	SignKey         *rsa.PrivateKey
	PublicCertPath  string `yaml:"public_cert"`
	VerifyKey       *rsa.PublicKey
}

func (jwtConfig *JwtConfig) InitJWT() *JwtConfig {
	jwtConfig.AccessToken = *jwtConfig.AccessToken.LoadKeys()
	jwtConfig.AccessToken.Duration = AccessDuration
	jwtConfig.RefreshToken = *jwtConfig.RefreshToken.LoadKeys()
	jwtConfig.RefreshToken.Duration = RefreshDuration
	return jwtConfig
}

func (jwtSpecs *JwtSpecs) LoadKeys() *JwtSpecs {
	jwtSpecs.SignKey = LoadPrivateKey(jwtSpecs.PrivateCertPath)
	jwtSpecs.VerifyKey = LoadPublicKey(jwtSpecs.PublicCertPath)
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
