package routers

import (
	"github.com/gin-gonic/gin"
	"goweb/controllers"
	"goweb/logger"
	"goweb/middleware"
	"goweb/settings"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("api/v1")
	//注册
	v1.POST("/signup", controllers.SignUpHandler)
	//登录
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middleware.JWTAuth())
	{
		v1.GET("/community", controllers.GetCommunityList)
		v1.GET("/community/:id", controllers.GetCommunityHandler)

		v1.POST("/post", controllers.PostArticleHandler)
		v1.GET("/post/:id", controllers.GetArticleHandler)
		v1.GET("/postlist", controllers.GetArticleListHandler)
		//根据不同的排序条件获取帖子列表
		v1.GET("/postlists", controllers.GetArticleLists)
		//投票
		v1.POST("/vote", controllers.PostVote)
	}

	v1.GET("/", func(c *gin.Context) {
		c.ShouldBind("version")
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"version": settings.Conf.Version,
		})
	})
	return r
}
