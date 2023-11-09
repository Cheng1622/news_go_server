package middleware

import (
	"time"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/jwt"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/gin-gonic/gin"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("auth-token")
		//token不存在
		if tokenStr == "" {
			response.Error(c, code.Unauthorized, code.Unauthorized.Msg())
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, err := jwt.ParseToken(tokenStr)
		if err != nil {
			response.Error(c, code.Unauthorized, code.Unauthorized.Msg())
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			response.Error(c, code.Unauthorized, code.Unauthorized.Msg())
			c.Abort() //阻止执行
			return
		}
		clog.Log.Infoln("userid:", tokenStruck.UserId)
		c.Set("userid", tokenStruck.UserId)

		c.Next()
	}
}
