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

	clog.Log.Debug("a:", a)
	clog.Log.Infoln("a:", a)
	clog.Log.Warnln("a:", a)
	clog.Log.Errorln("a:", a)
	clog.Log.DPanic("a:", a)
	b := false
	if b {
		// 程序会崩溃并打印调用栈信息
		clog.Log.Panicln("a:", a)
	}
	// 不打印，os.Exit(1)
	clog.Log.Fatalln("a:", a)
}
