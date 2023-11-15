package test

import (
	"errors"
	"testing"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
	"github.com/Cheng1622/news_go_server/pkg/redis"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

func sn() int64 {
	a, _ := snowflake.SF.GenerateID()
	return a
}
func pw() string {
	a := "123456"
	b := encrypt.NewGenPasswd(a)
	return b
}

// 初始化mysql数据
func TestData(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	mysql.InitMysql()
	redis.InitRedis()
	snowflake.InitSnowflake()
	casbin.InitCasbinEnforcer()
	mysql.DB.AutoMigrate(&model.Api{}, &model.Role{}, &model.Menu{}, &model.User{})
	// 1.写入角色数据
	newRoles := make([]*model.Role, 0)
	roles := []*model.Role{
		{
			Name:    "超级管理员",
			Keyword: "admin",
			Desc:    "超级权限",
			Status:  1,
			Sort:    1,
			Creator: "admin",
		},
		{
			Name:    "普通管理员",
			Keyword: "edit",
			Desc:    "普通权限",
			Status:  1,
			Sort:    2,
			Creator: "admin",
		},
		{
			Name:    "用户",
			Keyword: "user",
			Desc:    "用户",
			Status:  1,
			Sort:    3,
			Creator: "admin",
		},
	}

	for _, role := range roles {
		err := mysql.DB.First(&role).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRoles = append(newRoles, role)
		}
	}

	if len(newRoles) > 0 {
		err := mysql.DB.Create(&newRoles).Error
		if err != nil {
			clog.Log.Errorln("写入admin角色数据失败:", err)
		}
	}

	// 3.写入用户
	newUsers := make([]*model.User, 0)
	a := "admin"
	b := "超级管理员"
	users := []*model.User{
		{
			Userid:       sn(),
			Username:     "admin",
			Password:     pw(),
			Mobile:       "12345678910",
			Nickname:     &a,
			Introduction: &b,
			Status:       1,
			Creator:      "admin",
			Roles:        roles[:1],
		},
	}

	for _, user := range users {
		err := mysql.DB.First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) > 0 {
		err := mysql.DB.Create(&newUsers).Error
		if err != nil {
			clog.Log.Errorln("写入用户数据失败:", err)
		}
	}

	// 4.写入api
	apis := []model.Api{
		{
			Method:   "POST",
			Path:     "api/v1/base/login",
			Category: "base",
			Desc:     "用户登录",
			Creator:  "admin",
		},
		{
			Method:   "POST",
			Path:     "api/v1/base/logout",
			Category: "base",
			Desc:     "用户登出",
			Creator:  "admin",
		},
		{
			Method:   "POST",
			Path:     "api/v1/base/refreshToken",
			Category: "base",
			Desc:     "刷新JWT令牌",
			Creator:  "admin",
		},
		{
			Method:   "POST",
			Path:     "api/v1/base/sendcode",
			Category: "base",
			Desc:     "给用户邮箱发送验证码",
			Creator:  "admin",
		},
		{
			Method:   "POST",
			Path:     "api/v1/base/changePwd",
			Category: "base",
			Desc:     "通过邮箱修改密码",
			Creator:  "admin",
		},
		{
			Method:   "GET",
			Path:     "api/v1/user/info",
			Category: "user",
			Desc:     "获取当前登录用户信息",
			Creator:  "admin",
		},
		{
			Method:   "GET",
			Path:     "api/v1/user/list",
			Category: "user",
			Desc:     "获取用户列表",
			Creator:  "admin",
		},
		{
			Method:   "PUT",
			Path:     "api/v1/user/changePwd",
			Category: "user",
			Desc:     "更新用户登录密码",
			Creator:  "admin",
		},
		{
			Method:   "POST",
			Path:     "api/v1/user/create",
			Category: "user",
			Desc:     "创建用户",
			Creator:  "admin",
		},
	}

	// 5. 将角色绑定给菜单
	newApi := make([]model.Api, 0)
	newRoleCasbin := make([]model.RoleCasbin, 0)
	for i, api := range apis {
		api.ID = uint(i + 1)
		err := mysql.DB.First(&api, api.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newApi = append(newApi, api)

			// 管理员拥有所有API权限
			newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})

			// 非管理员拥有基础权限
			basePaths := []string{
				"api/v1/base/login",
				"api/v1/base/captcha",
				"api/v1/user/info",
				"api/v1/user/changePwd",
			}

			if funk.ContainsString(basePaths, api.Path) {
				newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}

	if len(newApi) > 0 {
		if err := mysql.DB.Create(&newApi).Error; err != nil {
			clog.Log.Errorln("写入api数据失败:", err)
		}
	}

	if len(newRoleCasbin) > 0 {
		rules := make([][]string, 0)
		for _, c := range newRoleCasbin {
			rules = append(rules, []string{
				c.Keyword, c.Path, c.Method,
			})
		}
		isAdd, err := casbin.CasbinEnforcer.AddPolicies(rules)
		if !isAdd {
			clog.Log.Errorln("写入casbin数据失败:", err)
		}
	}

}
