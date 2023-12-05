package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
)

func TestEncrypt(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	a := encrypt.NewGenPasswd("123456")
	// b := encrypt.NewParPasswd(a)
	t.Fatal(a)
}
