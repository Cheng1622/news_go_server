package middleware

import (
	"strings"
	"sync"

	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware Casbin中间件, 基于RBAC的权限访问控制模型
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ur := service.NewUserService()
		user, err := ur.GetCurrentUser(c)
		if err != nil {
			response.Error(c, code.WithoutLogin, nil)
			c.Abort()
			return
		}
		if user.Status != 1 {
			response.Error(c, code.DisableAuth, nil)
			c.Abort()
			return
		}
		// 获得用户的全部角色
		roles := user.Roles
		// 获得用户全部未被禁用的角色的Keyword
		var subs []string
		for _, role := range roles {
			if role.Status == 1 {
				subs = append(subs, role.Keyword)
			}
		}
		// 获得请求路径URL
		obj := strings.TrimPrefix(c.FullPath(), "/")
		// 获取请求方式
		act := c.Request.Method
		isPass := Check(subs, obj, act)
		if !isPass {
			response.Error(c, code.AuthError, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

func Check(subs []string, obj string, act string) bool {
	// 同一时间只允许一个请求执行校验, 否则可能会校验失败
	var checkLock sync.Mutex
	checkLock.Lock()
	defer checkLock.Unlock()
	isPass := false
	for _, sub := range subs {
		pass, _ := casbin.CasbinEnforcer.Enforce(sub, obj, act)
		clog.Log.Errorln("aaa:", sub, obj, act)
		if pass {
			isPass = true
			break
		}
	}
	return isPass
}
