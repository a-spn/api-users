package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type MySQLConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

var (
	ErrorWithDatabase = errors.New("issue with the database")
	Db                *gorm.DB
)

func (mysqlConfig *MySQLConfig) InitDbConnection() {
	if mysqlConfig.Password == "" {
		mysqlConfig.Password = os.Getenv("MYSQL_PASSWORD")
		if mysqlConfig.Password == "" {
			Logger.Fatal("mysql database password is not set")
		}
	}
	var err error
	gormLogger := zapgorm2.New(Logger)
	gormLogger.SetAsDefault()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, strconv.Itoa(mysqlConfig.Port), mysqlConfig.Database)
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{Logger: gormLogger})
	if err != nil {
		Logger.Fatal("failed to connect to mysql database", zap.String("Username", mysqlConfig.Username), zap.String("host", mysqlConfig.Host), zap.Int("port", mysqlConfig.Port), zap.String("database", mysqlConfig.Database), zap.Error(err))
	}
	Logger.Info("successfuly connected to the database", zap.String("host", mysqlConfig.Host), zap.String("database", mysqlConfig.Database))
}
