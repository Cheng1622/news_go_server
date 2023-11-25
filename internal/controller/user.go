package controller

import (
	"strconv"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/repo"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
	"github.com/Cheng1622/news_go_server/pkg/jwt"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
	"github.com/Cheng1622/news_go_server/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type UserApi interface {
	Login(c *gin.Context)                // 登录
	Logout(c *gin.Context)               //注销
	RefreshToken(c *gin.Context)         //刷新token
	GetUserInfo(c *gin.Context)          // 获取当前登录用户信息
	GetUsers(c *gin.Context)             // 获取用户列表
	ChangePwd(c *gin.Context)            // 更新用户登录密码
	CreateUser(c *gin.Context)           // 创建用户
	UpdateUserById(c *gin.Context)       // 更新用户
	BatchDeleteUserByIds(c *gin.Context) // 批量删除用户
}

// UserApiService UserService 服务层数据处理
type UserApiService struct {
	User service.UserService
}

// NewUserApi 构造函数
func NewUserApi() UserApi {
	return UserApiService{User: service.NewUserService()}
}

// Login 校验token的正确性, 处理登录逻辑
//
//	@Tags		Base
//	@Summary	用户登录
//	@Produce	application/json
//	@Param		data	body		requ.RegisterAndLoginRequest	true	"用户名,密码,验证码"
//	@Success	1000	{object}	response.Response{data}			"用户登录,返回token"
//	@Router		/api/v1/base/login [post]
func (us UserApiService) Login(c *gin.Context) {
	var req requ.RegisterAndLoginRequest
	// 请求json绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 密码校验
	user, err := us.User.Login(&req)
	if err != nil {
		response.Error(c, code.LoginError, err.Error())
		return
	}
	token, err := jwt.GenToken(*user)
	if err != nil {
		response.Error(c, code.LoginError, err.Error())
		return
	}
	response.Success(c, code.SUCCESS,
		gin.H{
			"token": token,
		})
}

// Logout 用户注销
func (us UserApiService) Logout(c *gin.Context) {
	err := jwt.JoinBlackList(c)
	if err != nil {
		response.Error(c, code.LogoutError, err.Error())
		return
	}
	response.Success(c, code.SUCCESS, nil)
}

// RefreshToken token续约
func (us UserApiService) RefreshToken(c *gin.Context) {

}

// GetUserInfo 获取当前登录用户信息
//
//	@Tags		User
//	@Summary	获取当前登录用户信息
//	@Produce	application/json
//	@Success	1000	{object}	response.Response{data=repo.UserInfoResp}	"获取当前登录用户信息,返回userInfo"
//	@Router		/api/v1/user/info [get]
func (us UserApiService) GetUserInfo(c *gin.Context) {
	user, err := us.User.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}

	userInforesponsep := repo.ToUserInfoResp(user)
	// 成功返回
	response.Success(c, code.SUCCESS, gin.H{
		"userInfo": userInforesponsep,
	})
}

// GetUsers 获取用户列表
//
//	@Tags		User
//	@Summary	获取用户列表
//	@Produce	application/json
//	@Success	1000	{object}	response.Response{data=[]repo.UserInfoResp}	"获取用户列表,返回userInfo和total"
//	@Router		/api/v1/user/list [get]
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
	response.Success(c, code.SUCCESS, gin.H{
		"users": repo.ToUsersResp(users),
		"total": total,
	})
}

// ChangePwd 更新用户登录密码
//
//	@Tags		User
//	@Summary	更新用户登录密码
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		requ.ChangePwdRequest		true	"旧密码,新密码"
//	@Success	1000	{object}	response.Response{}	"更新用户登录密码,返回成功"
//	@Router		/api/v1/user/changePwd [put]
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
		response.Error(c, code.ServerErr, err.Error())
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
}

// CreateUser 创建用户
//
//	@Tags		User
//	@Summary	创建用户
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		requ.CreateUserRequest		true	"用户信息"
//	@Success	1000	{object}	response.Response{}	"创建用户,返回成功"
//	@Router		/api/v1/user/create [post]
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
		clog.Log.Errorln("获取当前用户角色排序最小值失败:", err.Error())
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
		clog.Log.Errorln("未获取到角色信息")
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
		Userid:       userid,
		Username:     req.Username,
		Password:     encrypt.NewGenPasswd(req.Password),
		Mobile:       req.Mobile,
		Email:        req.Email,
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
		clog.Log.Errorln("创建用户失败:", err)
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
}

// UpdateUserById 更新用户
//
//	@Tags		User
//	@Summary	更新用户
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		requ.CreateUserRequest		true	"用户信息"
//	@Success	1000	{object}	response.Response{}	"更新用户,返回成功"
//	@Router		/api/v1/update/:userId [patch]
func (us UserApiService) UpdateUserById(c *gin.Context) {
	var req requ.CreateUserRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	//获取path中的userId
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userId <= 0 {
		// 错误返回
		response.Error(c, code.ServerErr, "用户ID不正确")
		return
	}

	// 根据path中的userId获取用户信息
	oldUser, err := us.User.GetUserById(userId)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取需要更新的用户信息失败: "+err.Error())
		return
	}

	// 获取当前用户
	ctxUser, err := us.User.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取当前用户失败: "+err.Error())
		return
	}
	// 获取当前用户的所有角色
	currentRoles := ctxUser.Roles
	// 获取当前用户角色的排序，和前端传来的角色排序做比较
	var currentRoleSorts []int
	// 当前用户角色ID集合
	var currentRoleIds []uint
	for _, role := range currentRoles {
		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
		currentRoleIds = append(currentRoleIds, role.ID)
	}
	// 当前用户角色排序最小值（最高等级角色）
	currentRoleSortMin := funk.MinInt(currentRoleSorts)

	// 获取前端传来的用户角色id
	reqRoleIds := req.RoleIds
	// 根据角色id获取角色
	rr := service.NewRoleService()
	roles, err := rr.GetRolesByIds(reqRoleIds)
	if err != nil {
		// 错误返回
		clog.Log.Errorln("根据角色ID获取角色信息失败: ", err)
		response.Error(c, code.ServerErr, "根据角色ID获取角色信息失败: "+err.Error())
		return
	}
	if len(roles) == 0 {
		// 错误返回
		clog.Log.Errorln("未获取到角色信息: ", err)
		response.Error(c, code.ServerErr, "未获取到角色信息: "+err.Error())
		return
	}
	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := funk.MinInt(reqRoleSorts)

	user := model.User{
		Model:        oldUser.Model,
		Userid:       oldUser.Userid,
		Username:     req.Username,
		Password:     oldUser.Password,
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Nickname:     &req.Nickname,
		Introduction: &req.Introduction,
		Status:       req.Status,
		Creator:      ctxUser.Username,
		Roles:        roles,
	}
	// 判断是更新自己还是更新别人
	if userId == ctxUser.Userid {
		// 如果是更新自己
		// 不能禁用自己
		if req.Status == 2 {
			// 错误返回
			clog.Log.Errorln("不能禁用自己")
			response.Error(c, code.ServerErr, "不能禁用自己")
			return
		}
		// 不能更改自己的角色
		reqDiff, currentDiff := funk.Difference(req.RoleIds, currentRoleIds)
		if len(reqDiff.([]uint)) > 0 || len(currentDiff.([]uint)) > 0 {
			// 错误返回
			clog.Log.Errorln("不能更改自己的角色")
			response.Error(c, code.ServerErr, "不能更改自己的角色")
			return
		}

		// 不能更新自己的密码，只能在个人中心更新
		if req.Password != "" {
			// 错误返回
			clog.Log.Errorln("请到个人中心更新自身密码")
			response.Error(c, code.ServerErr, "请到个人中心更新自身密码")
			return
		}

		// 密码赋值
		user.Password = ctxUser.Password

	} else {
		// 如果是更新别人
		// 用户不能更新比自己角色等级高的或者相同等级的用户
		// 根据path中的userIdID获取用户角色排序最小值
		minRoleSorts, err := us.User.GetUserMinRoleSortsByIds([]int64{userId})
		if err != nil || len(minRoleSorts) == 0 {
			// 错误返回
			response.Error(c, code.ServerErr, "根据用户ID获取用户角色排序最小值失败")
			return
		}
		if currentRoleSortMin >= minRoleSorts[0] {
			// 错误返回
			response.Error(c, code.ServerErr, "用户不能更新比自己角色等级高的或者相同等级的用户")
			return
		}

		// 用户不能把别的用户角色等级更新得比自己高或相等
		if currentRoleSortMin >= reqRoleSortMin {
			// 错误返回
			response.Error(c, code.ServerErr, "用户不能把别的用户角色等级更新得比自己高或相等")
			return
		}

		// 密码赋值
		if req.Password != "" {
			user.Password = encrypt.NewGenPasswd(req.Password)
		}
	}

	// 更新用户
	err = us.User.UpdateUser(&user)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "更新用户失败")
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
}

// BatchDeleteUserByIds 批量删除用户
//
//	@Tags		User
//	@Summary	批量删除用户
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		requ.DeleteUserRequest		true	"用户信息"
//	@Success	1000	{object}	response.Response{}	"批量删除用户,返回成功"
//	@Router		/api/v1/delete/batch [delete]
func (us UserApiService) BatchDeleteUserByIds(c *gin.Context) {
	var req requ.DeleteUserRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 前端传来的用户ID
	reqUserIds := req.Userids
	// 根据用户ID获取用户角色排序最小值
	roleMinSortList, err := us.User.GetUserMinRoleSortsByIds(reqUserIds)
	if err != nil || len(roleMinSortList) == 0 {
		// 错误返回
		response.Error(c, code.ServerErr, "根据用户ID获取用户角色排序最小值失败")
		return
	}

	// 当前用户角色排序最小值（最高等级角色）以及当前用户
	minSort, ctxUser, err := us.User.GetCurrentUserMinRoleSort(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	currentRoleSortMin := int(minSort)

	// 不能删除自己
	if funk.Contains(reqUserIds, ctxUser.ID) {
		// 错误返回
		response.Error(c, code.ServerErr, "用户不能删除自己")
		return
	}

	// 不能删除比自己角色排序低(等级高)的用户
	for _, sort := range roleMinSortList {
		if currentRoleSortMin >= sort {
			// 错误返回
			response.Error(c, code.ServerErr, "用户不能删除比自己角色等级高的用户")
			return
		}
	}

	err = us.User.BatchDeleteUserByIds(reqUserIds)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "删除用户失败: "+err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
}
