package router

import (
	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

// InitExampleRouter 用户案例模块
func InitDictRouter(r *gin.RouterGroup) gin.IRouter {
	dictApi := controller.NewDictApi()
	router := r.Group("/dict")
	// 开启jwt认证中间件
	router.Use(middleware.JwtMiddleware())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.POST("/create", dictApi.PostDict)        // 创建
		router.GET("/list", dictApi.GetDictList)        // 列表
		router.PUT("/put", dictApi.PutDict)             // 更新
		router.DELETE("/delete", dictApi.DeleteDict)    // 删除
		router.DELETE("/remove", dictApi.DeleteDictAll) // 全部删除
	}

	return r
}
