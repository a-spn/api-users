package config

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
)

const (
	RbacAuthorizationContextKey = "rbacClientData"
)

type RbacConfig struct {
	PathToCasbinModel  string `yaml:"model"`
	PathToCasbinPolicy string `yaml:"policy"`
	Enforcer           casbin.Enforcer
}

func (rbacConfig *RbacConfig) InitRBAC() (newConfig *RbacConfig) {
	enf, err := casbin.NewEnforcer(rbacConfig.PathToCasbinModel, rbacConfig.PathToCasbinPolicy)
	if err != nil {
		Logger.Fatal("can't initialize casbin Enforcer", zap.Error(err))
	}
	rbacConfig.Enforcer = *enf
	return rbacConfig
}
