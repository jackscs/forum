package controllers

import (
	"github.com/gin-gonic/gin"
	"goweb/logic"
	"strconv"
)

func GetCommunityList(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func GetCommunityHandler(c *gin.Context) {
	//获取URL中的参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	communityTable, err := logic.GetCommunityDetailList(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, communityTable)
}
