package router

import (
	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

// InitUserRouter 注册用户路由
func InitUserRouter(r *gin.RouterGroup) gin.IRouter {
	userApi := controller.NewUserApi()
	router := r.Group("/user")
	// 开启jwt认证中间件
	router.Use(middleware.JwtMiddleware())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/logout", userApi.Logout)
		router.POST("/refreshToken", userApi.RefreshToken)
		router.GET("/info", userApi.GetUserInfo)
		router.GET("/list", userApi.GetUsers)
		router.PUT("/changePwd", userApi.ChangePwd)
		router.POST("/create", userApi.CreateUser)
		router.PATCH("/update/:userId", userApi.UpdateUserById)
		router.DELETE("/delete/batch", userApi.BatchDeleteUserByIds)
	}
	return r
}
