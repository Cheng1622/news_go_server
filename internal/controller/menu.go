package controller

import (
	"strconv"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/Cheng1622/news_go_server/pkg/validator"
	"github.com/gin-gonic/gin"
)

type MenuApi interface {
	GetMenus(c *gin.Context)                // 获取菜单列表
	GetMenuTree(c *gin.Context)             // 获取菜单树
	CreateMenu(c *gin.Context)              // 创建菜单
	UpdateMenuById(c *gin.Context)          // 更新菜单
	BatchDeleteMenuByIds(c *gin.Context)    // 批量删除菜单
	GetUserMenusByUserId(c *gin.Context)    // 获取用户的可访问菜单列表
	GetUserMenuTreeByUserId(c *gin.Context) // 获取用户的可访问菜单树
}

type MenuApiService struct {
	Menu service.MenuService
}

func NewMenuMenuApi() MenuApi {
	return MenuApiService{Menu: service.NewMenuService()}
}

// 获取菜单列表
func (ms MenuApiService) GetMenus(c *gin.Context) {
	menus, err := ms.Menu.GetMenus()
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return

	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"menus": menus,
	})
	return
}

// 获取菜单树
func (ms MenuApiService) GetMenuTree(c *gin.Context) {
	menuTree, err := ms.Menu.GetMenuTree()
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"menuTree": menuTree,
	})
	return
}

// 创建菜单
func (ms MenuApiService) CreateMenu(c *gin.Context) {
	var req requ.CreateMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 获取当前用户
	ur := service.NewUserService()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取当前用户信息失败")
		return
	}

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = ms.Menu.CreateMenu(&menu)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "创建菜单失败: "+err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// 更新菜单
func (ms MenuApiService) UpdateMenuById(c *gin.Context) {
	var req requ.UpdateMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 获取路径中的menuId
	menuId, _ := strconv.Atoi(c.Param("menuId"))
	if menuId <= 0 {
		// 错误返回
		response.Error(c, code.ServerErr, "菜单ID不正确")
		return
	}

	// 获取当前用户
	ur := service.NewUserService()
	ctxUser, err := ur.GetCurrentUser(c)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取当前用户信息失败")
		return
	}

	menu := model.Menu{
		Name:       req.Name,
		Title:      req.Title,
		Icon:       &req.Icon,
		Path:       req.Path,
		Redirect:   &req.Redirect,
		Component:  req.Component,
		Sort:       req.Sort,
		Status:     req.Status,
		Hidden:     req.Hidden,
		NoCache:    req.NoCache,
		AlwaysShow: req.AlwaysShow,
		Breadcrumb: req.Breadcrumb,
		ActiveMenu: &req.ActiveMenu,
		ParentId:   &req.ParentId,
		Creator:    ctxUser.Username,
	}

	err = ms.Menu.UpdateMenuById(uint(menuId), &menu)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "更新菜单失败: "+err.Error())
		return
	}

	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return

}

// 批量删除菜单
func (ms MenuApiService) BatchDeleteMenuByIds(c *gin.Context) {
	var req requ.DeleteMenuRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}

	err := ms.Menu.BatchDeleteMenuByIds(req.MenuIds)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "删除菜单失败: "+err.Error())
		return
	}

	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// 根据用户ID获取用户的可访问菜单列表
func (ms MenuApiService) GetUserMenusByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		// 错误返回
		response.Error(c, code.ServerErr, "用户ID不正确 ")
		return
	}

	menus, err := ms.Menu.GetUserMenusByUserId(uint(userId))
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取用户的可访问菜单列表失败: "+err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"menus": menus,
	})
	return
}

// 根据用户ID获取用户的可访问菜单树
func (ms MenuApiService) GetUserMenuTreeByUserId(c *gin.Context) {
	// 获取路径中的userId
	userId, _ := strconv.Atoi(c.Param("userId"))
	if userId <= 0 {
		// 错误返回
		response.Error(c, code.ServerErr, "用户ID不正确")
		return
	}

	menuTree, err := ms.Menu.GetUserMenuTreeByUserId(uint(userId))
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, "获取用户的可访问菜单树失败: "+err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"menuTree": menuTree,
	})
	return
}
