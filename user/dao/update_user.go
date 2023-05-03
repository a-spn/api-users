package user_dao

import (
	"api-users/config"
	user_model "api-users/user/model"

	"go.uber.org/zap"
)

func (dao *UserDAO) UpdateUser(user user_model.User) (err error) {
	res := dao.db.Model(&user).Updates(user)
	if res.Error != nil {
		err = res.Error
		config.Logger.Error("error on updating user by id", zap.Error(err), zap.String("username", user.Username), zap.Uint("id", user.ID))
		return err
	}
	if res.RowsAffected == 0 {
		err = ErrUserDoesNotExist
	}
	return err
}
