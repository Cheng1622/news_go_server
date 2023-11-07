package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
)

func TestCasbin(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	casbin.InitCasbinEnforcer()

}
