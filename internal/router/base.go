package router

import (
	"net/http"

	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/gin-gonic/gin"
)

// InitBaseRouter 注册基础路由
func InitBaseRouter(r *gin.RouterGroup) gin.IRouter {
	baseApi := controller.NewCaptchaApi()
	userApi := controller.NewUserApi()
	router := r.Group("/base")
	{
		router.GET("ping", func(context *gin.Context) {
			context.String(http.StatusOK, "ok")
		})
		router.GET("/captcha", baseApi.Captcha)
		router.POST("/login", userApi.Login)
	}
	return r
}
