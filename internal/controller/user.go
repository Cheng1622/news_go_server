package controller

import (
	"encoding/json"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/repo"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
	"github.com/Cheng1622/news_go_server/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type UserApi interface {
	GetUserInfo(c *gin.Context) // 获取当前登录用户信息
	GetUsers(c *gin.Context)    // 获取用户列表
	ChangePwd(c *gin.Context)   // 更新用户登录密码
	CreateUser(c *gin.Context)  // 创建用户
}

// UserApiService UserService 服务层数据处理
type UserApiService struct {
	User service.UserService
}

// NewUserApi 构造函数
func NewUserApi() UserApi {
	return UserApiService{User: service.NewUserService()}
}

// GetUserInfo 获取当前登录用户信息
func (us UserApiService) GetUserInfo(c *gin.Context) {
	user, err := us.User.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	userInforesponsep := repo.ToUserInfoResp(user)
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"userInfo": userInforesponsep,
	})
	return

}

// GetUsers 获取用户列表
func (us UserApiService) GetUsers(c *gin.Context) {
	var req requ.UserListRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 获取
	users, total, err := us.User.GetUsers(&req)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"users": repo.ToUsersResp(users),
		"total": total,
	})
	return
}

// ChangePwd 更新用户登录密码
func (us UserApiService) ChangePwd(c *gin.Context) {
	var req requ.ChangePwdRequest

	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 获取当前用户
	user, err := us.User.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 获取用户的真实正确密码
	correctPasswd := user.Password
	// 判断前端请求的密码是否等于真实密码
	passwd := encrypt.NewParPasswd(correctPasswd)
	if passwd != req.OldPassword {
		// 错误返回
		response.Error(c, code.PasswordError, nil)
		return
	}
	// 更新密码
	err = us.User.ChangePwd(user.Username, encrypt.NewGenPasswd(req.NewPassword))
	if err != nil {
		// 错误返回
		response.Error(c, code.UpdateError, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// CreateUser 创建用户
func (us UserApiService) CreateUser(c *gin.Context) {
	var req requ.CreateUserRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 当前用户角色排序最小值（最高等级角色）以及当前用户
	currentRoleSortMin, ctxUser, err := us.User.GetCurrentUserMinRoleSort(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}

	// 获取前端传来的用户角色id
	reqRoleIds := req.RoleIds
	// 根据角色id获取角色
	rr := service.NewRoleService()
	roles, err := rr.GetRolesByIds(reqRoleIds)
	if err != nil {
		// 错误返回
		clog.Log.Errorln("根据角色ID获取角色信息失败: ", err)
		response.Error(c, code.ServerErr, nil)
		return
	}
	if len(roles) == 0 {
		// 错误返回
		clog.Log.Errorln("未获取到角色信息:")
		response.Error(c, code.ServerErr, nil)
		return
	}
	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := uint(funk.MinInt(reqRoleSorts))

	// 当前用户的角色排序最小值 需要小于 前端传来的角色排序最小值（用户不能创建比自己等级高的或者相同等级的用户）
	if currentRoleSortMin >= reqRoleSortMin {
		// 错误返回
		clog.Log.Errorln("用户不能创建比自己等级高的或者相同等级的用户")
		response.Error(c, code.ServerErr, nil)
		return
	}

	userid, _ := snowflake.SF.GenerateID()
	user := model.User{
		UserId:       userid,
		Username:     req.Username,
		Password:     encrypt.NewGenPasswd(req.Password),
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Nickname:     &req.Nickname,
		Introduction: &req.Introduction,
		Status:       req.Status,
		Creator:      ctxUser.Username,
		Roles:        roles,
	}

	err = us.User.CreateUser(&user)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return

}

// 校验token的正确性, 处理登录逻辑
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body  requ.RegisterAndLoginRequest true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data,msg=string} "返回包括用户信息,token,过期时间"
// @Router /api/base/login [post]
func Login(c *gin.Context) {
	var req requ.RegisterAndLoginRequest
	// 请求json绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 密码校验
	UserService := service.NewUserService()
	user, err := UserService.Login(&req)
	if err != nil {
		response.Error(c, code.AuthError, nil)
		return
	}
	userstr, err := json.Marshal(user)
	if err != nil {
		response.Error(c, code.ServerErr, nil)
		return
	}
	response.Success(c, code.SUCCESS,
		map[string]interface{}{
			"user": userstr,
		})

}
