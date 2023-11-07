package main

import (
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
	"github.com/Cheng1622/news_go_server/pkg/validator"
)

func init() {
	config.InitConfig()
	clog.InitLogger()
	snowflake.InitSnowflake()
	validator.InitValidate()
}
