package logic

import (
	"goweb/dao/mysql"
	"goweb/models"
	"goweb/pkg/jwt"
	"goweb/pkg/snowflake"
)

func SignUp(p *models.ParamsSignUp) error {
	//判断用户存不存在
	err := mysql.CheckUserNameExist(p.UserName)
	//fmt.Println(err)
	if err != nil {
		return err
	}
	//生成UID
	UserID := snowflake.GenID()
	User := &models.User{
		UserID:   UserID,
		Username: p.UserName,
		Password: p.Password,
	}

	//保存进数据库
	err = mysql.InsertUser(User)
	//if err != nil {
	//	return errors.New("添加用户失败")
	//}

	return err
}

func Logic(p *models.ParamsLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//通过传入指针，获取user.ID
	m, err := mysql.Login(user)
	if err != nil {
		return "", err
	}
	token, err = jwt.GenToken(m.UserID, user.Username)
	return token, nil
}
