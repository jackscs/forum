package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb/logic"
	"goweb/models"
	"strconv"
)

// PostArticleHandler 创建帖子的函数
func PostArticleHandler(c *gin.Context) {
	//获取参数并进行校验
	article := new(models.PostArticle)
	err := c.ShouldBindJSON(&article)
	if err != nil {
		ResponseError(c, CodeInvalidArticle)
		return
	}

	//从请求的token中获取作者ID
	article.AuthorID, err = GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeInvalidAuthorID)
	}
	//创建帖子
	err = logic.PostArticle(article)
	if err != nil {
		ResponseError(c, CodeInvalidArticle)
		return
	}
	//返回响应
	ResponseSuccess(c, CodeArticleSuccess)
}

// GetArticleHandler 获取帖子
func GetArticleHandler(c *gin.Context) {
	//获取URL中作品列表id的参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetArticleHandler 获取帖子id错误")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取帖子
	article, err := logic.GetArticle(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//向前端返回帖子
	ResponseSuccess(c, article)
}

// GetArticleListHandler 获取帖子列表
func GetArticleListHandler(c *gin.Context) {
	//获取页面大小和偏移量，传递给
	pagesizeStr := c.Query("pagesize")
	offsetStr := c.Query("offset")
	pagesize, err := strconv.ParseInt(pagesizeStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if pagesize == 0 {
		pagesize = 10
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if offset == 0 {
		offset = 5
	}

	data, err := logic.GetArticleList(pagesize, offset)
	if err != nil {
		zap.L().Error("GetArticleListHandler failed ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//根据时间或按照分数进行排序
//1.获取参数
//2.去redis查询id列表
//3.根据id查询数据库帖子详细信息
func GetArticleLists(c *gin.Context) {
	//按不同条件对帖子进行排序
	//初始化结构体初始化参数
	p := models.ParamArticleLists{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.L().Error("post_article.GetArticleLists failed")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := logic.GetArticleLists(&p)
	if err != nil {
		zap.L().Error("GetArticleLists 获取数据错误")
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
