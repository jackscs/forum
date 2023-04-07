package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/controllers"
	"goweb/pkg/jwt"
	"strings"
)

//从上下文中获取token并进行验证

func JWTAuth() func(c *gin.Context) {
	//token请求携带方式有三种：请求头里面，请求体里面，URL里面

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		//fmt.Println(authHeader)
		//authHeader : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlc3RpbnkiLCJleHAiOjE2NjE4NDk2NjMsImlzcyI6ImRlc3Rp
		//bnkiLCJuYmYiOjE2NjE4NDI0NjN9.xs-kMew1Z9tjK4kg8qZNFgnMVWpqxa81gMD8X8HF6cI
		//按空格分割authHeader
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		//将当前请求的userID信息保存到请求的上下文中
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		//fmt.Println(mc)

		//判断token是否和redis缓存中的token相同 实现单点登录
		//rtoken := redis.GetToken(claim.UserID)
		//if rtoken != parts[1] {
		//	controllers.ResponseError(c, controllers.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}
		c.Next() //后续的请求处理函数中可以通过c.Get(CtxUserIDKey)获取当前用户请求信息
	}
}
