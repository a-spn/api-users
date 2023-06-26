package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MySQL    MySQLConfig    `yaml:"mysql"`
	Api      ApiConfig      `yaml:"api"`
	JWT      JwtConfig      `yaml:"jwt"`
	Rbac     RbacConfig     `yaml:"rbac"`
	Security SecurityConfig `yaml:"security"`
}

var (
	Configuration Config
	ConfigPath    string
)

func ParseConfig() {
	yamlFile, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatalf("cant read file config file '%s' : %s", ConfigPath, err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Configuration)
	if err != nil {
		log.Fatalf("cant parse config file '%s' : %s", ConfigPath, err.Error())
	}
}

func InitConfig() {
	InitLogger()
	Configuration.Rbac = *Configuration.Rbac.InitRBAC()
	Configuration.JWT = *Configuration.JWT.InitJWT()
	Configuration.Security = *Configuration.Security.InitSecurity()

}
