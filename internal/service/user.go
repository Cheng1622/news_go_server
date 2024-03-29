package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/pkg/captcha"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/thoas/go-funk"
)

type UserService interface {
	Login(user *requ.RegisterAndLoginRequest) (*int64, error)           // 登录
	ChangePwd(username string, newPasswd string) error                  // 更新密码
	CreateUser(user *model.User) error                                  // 创建用户
	GetUserById(id int64) (model.User, error)                           // 获取单个用户
	GetUsers(req *requ.UserListRequest) ([]*model.User, int64, error)   // 获取用户列表
	UpdateUser(user *model.User) error                                  // 更新用户
	BatchDeleteUserByIds(ids []int64) error                             // 批量删除
	GetCurrentUser(c *gin.Context) (model.User, error)                  // 获取当前登录用户信息
	GetCurrentUserMinRoleSort(c *gin.Context) (uint, model.User, error) // 获取当前用户角色排序最小值（最高等级角色）以及当前用户信息
	GetUserMinRoleSortsByIds(ids []int64) ([]int, error)                // 根据用户ID获取用户角色排序最小值
	SetUserInfoCache(username string, user model.User)                  // 设置用户信息缓存
	UpdateUserInfoCacheByRoleId(roleId uint) error                      // 根据角色ID更新拥有该角色的用户信息缓存
	ClearUserInfoCache()                                                // 清理所有用户信息缓存
}

type User struct{}

// 当前用户信息缓存，避免频繁获取数据库
var userInfoCache = cache.New(24*time.Hour, 48*time.Hour)

// NewUserService 构造函数
func NewUserService() UserService {
	return User{}
}

// Login 登录
func (ud User) Login(user *requ.RegisterAndLoginRequest) (*int64, error) {
	// 根据用户名获取用户(正常状态:用户状态正常)
	var firstUser model.User
	err := mysql.DB.
		Where("username = ?", user.Username).
		Preload("Roles").
		First(&firstUser).Error

	clog.Log.Errorln(err)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 判断用户的状态
	userStatus := firstUser.Status
	if userStatus != 1 {
		return nil, errors.New("用户被禁用")
	}

	// 判断用户拥有的所有角色的状态,全部角色都被禁用则不能登录
	roles := firstUser.Roles
	isValidate := false
	for _, role := range roles {
		// 有一个正常状态的角色就可以登录
		if role.Status == 1 {
			isValidate = true
			break
		}
	}

	if !isValidate {
		return nil, errors.New("用户角色被禁用")
	}

	// 校验密码
	passwd := encrypt.NewParPasswd(firstUser.Password)
	if passwd != user.Password {
		return nil, errors.New("密码错误")
	}

	// 获取 redis数字验证码
	CaptchaCache := captcha.NewCaptchaService()
	code := CaptchaCache.GetCaptcha(user.CaptchaId)
	if code == "" {
		return nil, errors.New("验证码过期")
	}
	if user.Code != code {
		return nil, errors.New("验证码错误")
	}
	return &firstUser.Userid, nil
}

// GetCurrentUser 获取当前登录用户信息
// 需要缓存，减少数据库访问
func (ud User) GetCurrentUser(c *gin.Context) (model.User, error) {
	var user model.User
	ctxUser, ok := c.Get("userid")
	if !ok {
		return user, errors.New("用户未登录")
	}
	user.Userid = ctxUser.(int64)
	userid := strconv.FormatInt(user.Userid, 10)
	// 先获取缓存
	cacheUser, found := userInfoCache.Get(userid)
	var err error
	if found {
		user = cacheUser.(model.User)
		err = nil
	} else {
		// 缓存中没有就获取数据库
		user, err = ud.GetUserById(user.Userid)
		// 获取成功就缓存
		if err != nil {
			userInfoCache.Delete(userid)
		} else {
			userInfoCache.Set(userid, user, cache.DefaultExpiration)
		}
	}
	return user, err
}

// GetCurrentUserMinRoleSort 获取当前用户角色排序最小值（最高等级角色）以及当前用户信息
func (ud User) GetCurrentUserMinRoleSort(c *gin.Context) (uint, model.User, error) {
	// 获取当前用户
	ctxUser, err := ud.GetCurrentUser(c)
	if err != nil {
		return 999, ctxUser, err
	}
	// 获取当前用户的所有角色
	currentRoles := ctxUser.Roles
	// 获取当前用户角色的排序，和前端传来的角色排序做比较
	var currentRoleSorts []int
	for _, role := range currentRoles {
		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
	}
	// 当前用户角色排序最小值（最高等级角色）
	currentRoleSortMin := uint(funk.MinInt(currentRoleSorts))

	return currentRoleSortMin, ctxUser, nil
}

// GetUserById 获取单个用户
func (ud User) GetUserById(id int64) (model.User, error) {
	var user model.User
	err := mysql.DB.Where("userid = ?", id).Preload("Roles").First(&user).Error
	return user, err
}

// GetUsers 获取用户列表
func (ud User) GetUsers(req *requ.UserListRequest) ([]*model.User, int64, error) {
	var list []*model.User
	db := mysql.DB.Model(&model.User{}).Order("created_at DESC")

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", nickname))
	}
	mobile := strings.TrimSpace(req.Mobile)
	if mobile != "" {
		db = db.Where("mobile LIKE ?", fmt.Sprintf("%%%s%%", mobile))
	}
	status := req.Status
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Preload("Roles").Find(&list).Error
	} else {
		err = db.Preload("Roles").Find(&list).Error
	}
	return list, total, err
}

// ChangePwd 更新密码
func (ud User) ChangePwd(username string, hashNewPasswd string) error {
	err := mysql.DB.Model(&model.User{}).Where("username = ?", username).Update("password", hashNewPasswd).Error
	// 如果更新密码成功，则更新当前用户信息缓存
	// 先获取缓存
	cacheUser, found := userInfoCache.Get(username)
	if err == nil {
		if found {
			user := cacheUser.(model.User)
			user.Password = hashNewPasswd
			userInfoCache.Set(username, user, cache.DefaultExpiration)
		} else {
			// 没有缓存就获取用户信息缓存
			var user model.User
			mysql.DB.Where("username = ?", username).First(&user)
			userInfoCache.Set(username, user, cache.DefaultExpiration)
		}
	}

	return err
}

// CreateUser 创建用户
func (ud User) CreateUser(user *model.User) error {
	err := mysql.DB.Create(user).Error
	return err
}

// UpdateUser 更新用户
func (ud User) UpdateUser(user *model.User) error {
	err := mysql.DB.Model(user).Updates(user).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Model(user).Association("Roles").Replace(user.Roles)

	//err := mysql.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error

	// 如果更新成功就更新用户信息缓存
	if err == nil {
		userInfoCache.Set(user.Username, *user, cache.DefaultExpiration)
	}
	return err
}

// BatchDeleteUserByIds 批量删除
func (ud User) BatchDeleteUserByIds(ids []int64) error {
	// 用户和角色存在多对多关联关系
	var users []model.User
	for _, id := range ids {
		// 根据ID获取用户
		user, err := ud.GetUserById(id)
		if err != nil {
			return errors.New(fmt.Sprintf("未获取到ID为%d的用户", id))
		}
		users = append(users, user)
	}

	err := mysql.DB.Select("Roles").Unscoped().Delete(&users).Error
	// 删除用户成功，则删除用户信息缓存
	if err == nil {
		for _, user := range users {
			userInfoCache.Delete(user.Username)
		}
	}
	return err
}

// GetUserMinRoleSortsByIds 根据用户ID获取用户角色排序最小值
func (ud User) GetUserMinRoleSortsByIds(ids []int64) ([]int, error) {
	// 根据用户ID获取用户信息
	var userList []model.User
	err := mysql.DB.Where("id IN (?)", ids).Preload("Roles").Find(&userList).Error
	if err != nil {
		return []int{}, err
	}
	if len(userList) == 0 {
		return []int{}, errors.New("未获取到任何用户信息")
	}
	var roleMinSortList []int
	for _, user := range userList {
		roles := user.Roles
		var roleSortList []int
		for _, role := range roles {
			roleSortList = append(roleSortList, int(role.Sort))
		}
		roleMinSort := funk.MinInt(roleSortList)
		roleMinSortList = append(roleMinSortList, roleMinSort)
	}
	return roleMinSortList, nil
}

// SetUserInfoCache 设置用户信息缓存
func (ud User) SetUserInfoCache(username string, user model.User) {
	userInfoCache.Set(username, user, cache.DefaultExpiration)
}

// UpdateUserInfoCacheByRoleId 根据角色ID更新拥有该角色的用户信息缓存
func (ud User) UpdateUserInfoCacheByRoleId(roleId uint) error {

	var role model.Role
	err := mysql.DB.Where("id = ?", roleId).Preload("Users").First(&role).Error
	if err != nil {
		return errors.New("根据角色ID角色信息失败")
	}

	users := role.Users
	if len(users) == 0 {
		return errors.New("根据角色ID未获取到拥有该角色的用户")
	}

	// 更新用户信息缓存
	for _, user := range users {
		_, found := userInfoCache.Get(user.Username)
		if found {
			userInfoCache.Set(user.Username, *user, cache.DefaultExpiration)
		}
	}

	return err
}

// ClearUserInfoCache 清理所有用户信息缓存
func (ud User) ClearUserInfoCache() {
	userInfoCache.Flush()
}
