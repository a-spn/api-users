package user_dao

import (
	"api-users/config"
	user_model "api-users/user/model"

	"go.uber.org/zap"
)

func (dao *UserDAO) GetAPIUsers(limit, offset int, visibleRoles []string) (users []user_model.APIUser, err error) {
	err = dao.db.Model(&user_model.User{}).Limit(limit).Offset(offset).Where("role IN ?", visibleRoles).Find(&users).Error
	if err != nil {
		config.Logger.Error("error on getting users", zap.Error(err))
	}
	return users, err
}

func (dao *UserDAO) GetAPIUsersUnscopped(limit, offset int, visibleRoles []string) (users []user_model.APIUser, err error) {
	err = dao.db.Unscoped().Model(&user_model.User{}).Limit(limit).Offset(offset).Where("role IN ?", visibleRoles).Find(&users).Error
	if err != nil {
		config.Logger.Error("error on getting users", zap.Error(err))
	}
	return users, err
}
