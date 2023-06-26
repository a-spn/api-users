package user_dao

import (
	"api-users/config"
	user_model "api-users/user/model"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func (dao *UserDAO) CreateUser(user user_model.User) (err error) {
	err = dao.db.Create(&user).Error
	if err != nil {
		config.Logger.Error("error on inserting a new user in db", zap.Error(err), zap.String("username", user.Username))
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrDuplicateKeyEntry
			}
		}
	}
	return err
}
