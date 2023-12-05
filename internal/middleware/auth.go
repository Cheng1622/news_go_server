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
func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("auth-token")
		//token不存在
		if tokenStr == "" {
			clog.Log.Errorln("token不存在")
			response.Error(c, code.Unauthorized, "token不存在")
			c.Abort() //阻止执行
			return
		}
		// 黑名单
		if jwt.IsInBlacklist(tokenStr) {
			response.Error(c, code.LoginNoError, "您的帐户异地登陆或令牌失效")
			c.Abort()
			return
		}
		//验证token
		tokenStruck, err := jwt.ParseToken(tokenStr)
		if err != nil {
			clog.Log.Errorln("token错误:", err)
			response.Error(c, code.Unauthorized, "token错误: "+err.Error())
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			clog.Log.Errorln("token超时")
			response.Error(c, code.Unauthorized, "token超时")
			c.Abort() //阻止执行
			return
		}

		c.Set("userid", tokenStruck.Userid)
		c.Set("expiresAt", tokenStruck.ExpiresAt)
		c.Set("tokenStr", tokenStr)

		c.Next()
	}
}
