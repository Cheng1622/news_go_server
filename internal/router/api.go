package router

import (
	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(r *gin.RouterGroup) gin.IRouter {
	apiApi := controller.NewApiApi()
	router := r.Group("/api")
	// 开启jwt认证中间件
	router.Use(middleware.JwtMiddleware())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())

	{
		router.GET("/list", apiApi.GetApis)
		router.GET("/tree", apiApi.GetApiTree)
		router.POST("/create", apiApi.CreateApi)
		router.PATCH("/update/:apiId", apiApi.UpdateApiById)
		router.DELETE("/delete/batch", apiApi.BatchDeleteApiByIds)
	}

	return r
}
