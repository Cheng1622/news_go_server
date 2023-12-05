package main

import (
	"github.com/Cheng1622/news_go_server/internal/server"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
	"github.com/Cheng1622/news_go_server/pkg/redis"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
	"github.com/Cheng1622/news_go_server/pkg/validator"
)

// @title			news_go_server
// @version		v.1.0.0
// @description news_go_server业务接口文档集合
// @contact.name	Cheng1622
// @contact.email	cchen1622@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	config.InitConfig()
	clog.InitLogger()
	mysql.InitMysql()
	redis.InitRedis()
	snowflake.InitSnowflake()
	casbin.InitCasbinEnforcer()
	validator.InitValidate()
	server.InitRun()
}
