package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
)

func TestMysql(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	mysql.InitMysql()
}
