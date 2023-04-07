package mysql

import (
	"go.uber.org/zap"
	"goweb/global"
	"goweb/models"
	"goweb/pkg/md5"
)

func CheckUserNameExist(username string) (err error) {
	var sum int64
	err = global.DBEngine.Table("user").Where("username=?", username).Count(&sum).Error
	if err != nil {
		return err
	}
	if sum > 0 {
		return ErrUserExist
	}
	return nil
}

func InsertUser(user *models.User) error {
	//对用户密码进进行加密
	user.Password = md5.EncryptPassword(user.Password)
	//操作数据库进行添加用户
	err := global.DBEngine.Model(&user).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(user *models.User) (users *models.User, err error) {

	var m models.User
	err = global.DBEngine.Table("user").Where("username=?", user.Username).Find(&m).Error
	if err != nil {
		return nil, ErrUserNotExit
	}
	password := md5.EncryptPassword(user.Password)
	if user.Username != m.Username || password != m.Password {
		return nil, ErrInvalidPassword
	}

	return &m, nil
}

//通过用户id查询用户信息
func GetUserByID(id int64) (*models.User, error) {
	data := new(models.User)
	err := global.DBEngine.Table("user").Where("user_id=?", id).Find(data).Error
	if err != nil {
		zap.L().Error("GetUserByID(id int64) failed")
		return nil, err
	}
	return data, nil
}
