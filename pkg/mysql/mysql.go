package mysql

import (
	"fmt"
	"time"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // DB 全局mysql数据库变量

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 执行任何 SQL 时都会创建一个 prepared statement 并将其缓存
		PrepareStmt: false,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		// 禁用嵌套事务
		DisableNestedTransaction: true,
	})
	if err != nil {
		clog.Log.Fatalln("初始化mysql数据库异常:", err)
	}
	if config.Conf.Mysql.LogMode {
		DB.Debug()
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		clog.Log.Fatalln("初始化mysql.DB数据库异常:", err)
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	clog.Log.Infoln("初始化mysql数据库完成!")

}
