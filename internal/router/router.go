package router

import (
	"time"

	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/gin-gonic/gin"
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

}
