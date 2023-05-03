package user_dao

import (
	"api-users/config"
	user_model "api-users/user/model"

	"go.uber.org/zap"
)

func (dao *UserDAO) GetAPIUserByIDUnscopped(id uint) (user *user_model.APIUser, err error) {
	err = dao.db.Unscoped().Model(&user_model.User{}).First(&user, id).Error
	if err != nil {
		config.Logger.Error("error on getting user by id", zap.Error(err), zap.Uint("id", id))
		return &user_model.APIUser{}, err
	}
	return user, nil
}

func (dao *UserDAO) GetUserByUsername(username string) (user *user_model.User, err error) {
	err = dao.db.First(&user, "username = ?", username).Error
	if err != nil {
		config.Logger.Error("error on getting user by username", zap.Error(err), zap.String("username", username))
		return &user_model.User{}, err
	}
	return user, nil
}

func (dao *UserDAO) GetUserByEmail(email string) (user *user_model.User, err error) {
	err = dao.db.Where(user, "email = ?", email).Find(&user).Error
	if err != nil {
		config.Logger.Error("error on getting user by email", zap.Error(err), zap.String("email", email))
		return &user_model.User{}, err
	}
	return user, nil
}

func (dao *UserDAO) GetUserByID(id uint) (user *user_model.User, err error) {
	err = dao.db.Model(&user_model.User{}).First(&user, id).Error
	if err != nil {
		config.Logger.Error("error on getting user by id", zap.Error(err), zap.Uint("id", id))
		return &user_model.User{}, err
	}
	return user, nil
}
