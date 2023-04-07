package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb/dao/mysql"
	"goweb/logic"
	"goweb/models"
)

func SignUpHandler(c *gin.Context) {
	//获取参数的数据
	var p models.ParamsSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signUp with param", zap.Error(err))
		//errs, ok := err.(validator.ValidationErrors)
		//if !ok {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": err.Error(),
		//	})
		//	return
		//}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
		return
	}

	//业务逻辑处理
	err := logic.SignUp(&p)
	if err != nil {
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回参数
	ResponseSuccess(c, "注册成功")
}

func LoginHandler(c *gin.Context) {
	//获取参数
	var p models.ParamsLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login failed", zap.Error(err))
	}
	//逻辑处理
	token, err := logic.Logic(&p)
	if err != nil {
		switch err {
		case mysql.ErrInvalidPassword:
			ResponseError(c, CodeInvalidPassword)
		case mysql.ErrUserNotExit:
			ResponseError(c, CodeUserNotExit)
		default:
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	//返回参数
	ResponseSuccess(c, token)
}
