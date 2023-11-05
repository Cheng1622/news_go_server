package test

import (
	"fmt"
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
)

func TestSnowFlake(t *testing.T) {
	config.InitConfig()
	snowflake.InitSnowflake()

	a, err := snowflake.SF.GenerateID()
	if err != nil {
		clog.Log.Panicln("snowflake:", err)
		return
	}
	fmt.Println(a)
}
