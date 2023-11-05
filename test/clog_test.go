package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
)

func Test(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	a := 1
	clog.Log.DPanic("a:", a)
}
