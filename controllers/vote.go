package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"goweb/logic"
	"goweb/models"
)

func PostVote(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(&p); err != nil {
		_, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		//errData := romove(errs.Translate(trans)) //翻译并去除错误结构体中标识
		//ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		//return
	}
	//获取当前请求用户的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//具体投票业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
