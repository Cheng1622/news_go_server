package router

import (
	"github.com/Cheng1622/news_go_server/internal/controller"
	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(r *gin.RouterGroup) gin.IRouter {
	roleApi := controller.NewRoleApi()
	router := r.Group("/role")
	// 开启jwt认证中间件
	router.Use(middleware.JwtMiddleware())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/list", roleApi.GetRoles)
		router.POST("/create", roleApi.CreateRole)
		router.PATCH("/update/:roleId", roleApi.UpdateRoleById)
		router.GET("/menus/get/:roleId", roleApi.GetRoleMenusById)
		router.PATCH("/menus/update/:roleId", roleApi.UpdateRoleMenusById)
		router.GET("/apis/get/:roleId", roleApi.GetRoleApisById)
		router.PATCH("/apis/update/:roleId", roleApi.UpdateRoleApisById)
		router.DELETE("/delete/batch", roleApi.BatchDeleteRoleByIds)
	}
	return r
}
