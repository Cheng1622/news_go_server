package router

import (
	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/gin-gonic/gin"
)

// InitUserRouter 注册用户路由
func InitUploadRouter(r *gin.RouterGroup) gin.IRouter {
	uploadApi := controller.NewUploadApi()
	router := r.Group("/upload")

	// 开启jwt认证中间件
	// router.Use(middleware.JwtMiddleware())
	// 开启casbin鉴权中间件
	// router.Use(middleware.CasbinMiddleware())
	{
		router.POST("/image", uploadApi.UploadImage) // 上传图片
	}
	return r
}
