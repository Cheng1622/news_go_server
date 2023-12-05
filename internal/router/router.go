package router

import (
	"time"

	_ "github.com/Cheng1622/news_go_server/docs"
	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化
func InitRouter() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)
	r := gin.Default()

	// 启用限流中间件
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity

	// 每fillInterval毫秒产生1个令牌，桶容量是capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	GroupRouter(r)
	// r.Static(relativePath string, root string)
	clog.Log.Infoln("初始化路由完成！")

	return r
}

func GroupRouter(r *gin.Engine) {
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	apiGroup := r.Group("/api/v1")
	// 注册路由
	InitBaseRouter(apiGroup) // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	InitUserRouter(apiGroup) // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	InitRoleRouter(apiGroup) // 注册角色路由, jwt认证中间件,casbin鉴权中间件
}
