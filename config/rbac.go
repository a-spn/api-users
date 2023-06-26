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

	Enforcer *casbin.Enforcer
}

func (rbacConfig *RbacConfig) InitRBAC() (newConfig *RbacConfig) {
	enforcer, err := casbin.NewEnforcer(rbacConfig.PathToCasbinModel, rbacConfig.PathToCasbinPolicy)
	if err != nil {
		Logger.Fatal("can't initialize casbin Enforcer", zap.Error(err))
	}
	rbacConfig.Enforcer = enforcer
	return rbacConfig
}

func (rbacConfig *RbacConfig) ListAllRoles() (roles []string) {
	return rbacConfig.Enforcer.GetAllObjects()
}

func (rbacConfig *RbacConfig) IsRoleValid(role string) bool {
	for _, r := range rbacConfig.ListAllRoles() {
		if role == r {
			return true
		}
	}
	return false
}
