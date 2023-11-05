package casbin

import (
	"os"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/mysql"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 全局CasbinEnforcer
var CasbinEnforcer *casbin.Enforcer

// 初始化casbin策略管理器
func InitCasbinEnforcer() {
	var err error
	CasbinEnforcer, err = mysqlCasbin()
	if err != nil {
		clog.Log.DPanicln("初始化Casbin失败:", err)
		os.Exit(1)
	}

	clog.Log.Infoln("初始化Casbin完成!")
}

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(config.Conf.Casbin.ModelPath, a)
	if err != nil {
		return nil, err
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
