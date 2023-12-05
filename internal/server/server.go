package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Cheng1622/news_go_server/internal/router"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
)

// InitRun 开启服务
func InitRun() {
	// 注册所有路由
	r := router.InitRouter()

	port := config.Conf.System.Port

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			clog.Log.Fatalln("监听端口失败: %s\n", err)
		}
	}()

	clog.Log.Infoln("服务启动完成")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	clog.Log.Infoln("正在关闭服务...")

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		clog.Log.Fatalln("服务关闭问题:", err)
	}

	clog.Log.Infoln("服务关闭成功!")
}
