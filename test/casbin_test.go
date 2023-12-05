package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/internal/middleware"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
)

func TestCasbin(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	mysql.InitMysql()
	casbin.InitCasbinEnforcer()
	middleware.Check([]string{"admin"}, "api/v1/user/info", "GET")
}
