package user_dao

import (
	"api-users/config"
	user_model "api-users/user/model"

	"go.uber.org/zap"
)

func (dao *UserDAO) DeleteUser(id uint) (err error) {
	err = dao.db.Delete(&user_model.User{}, id).Error
	if err != nil {
		config.Logger.Error("failed to delete user", zap.Error(err))
	}
	return err
}
