package test

import (
	"fmt"
	"testing"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
	"github.com/Cheng1622/news_go_server/pkg/redis"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
)

func init() {
	config.InitConfig()
	clog.InitLogger()
	mysql.InitMysql()
	redis.InitRedis()
	snowflake.InitSnowflake()
	casbin.InitCasbinEnforcer()
}

// 获取菜单列表
func TestGetMenus(t *testing.T) {
	menu, err := service.NewMenuService().GetMenus()
	if err != nil {
		t.Fatal("获取菜单列表失败: ", err)
	}
	for _, v := range menu {
		fmt.Println(*v)
	}
}

// 获取菜单树
func TestGetMenuTree(t *testing.T) {
	menu, err := service.NewMenuService().GetMenuTree()
	if err != nil {
		t.Fatal("获取菜单树失败: ", err)
	}
	for _, v := range menu {
		fmt.Println(*v)
	}
}

// 创建菜单
func TestCreateMenu(t *testing.T) {
	icon := "tree"
	parentid := uint(0)
	req := &model.Menu{
		Name:       "权限管理",
		Title:      "权限管理",
		Icon:       &icon,
		Path:       "/permission",
		Component:  "Layout",
		Sort:       2,
		Status:     1,
		Hidden:     2,
		NoCache:    2,
		AlwaysShow: 2,
		Breadcrumb: 1,
		ParentId:   &parentid,
		Creator:    "admin",
	}
	err := service.NewMenuService().CreateMenu(req)
	if err != nil {
		t.Fatal("创建菜单失败: ", err)
	}
	fmt.Println("成功")
}

// 更新菜单
func TestUpdateMenuById(t *testing.T) {
	req := &model.Menu{
		Name: "权限管理",
	}
	err := service.NewMenuService().UpdateMenuById(1, req)
	if err != nil {
		t.Fatal("更新菜单失败: ", err)
	}
	fmt.Println("成功")
}

// 批量删除菜单
func TestBatchDeleteMenuByIds(t *testing.T) {
	req := []uint{1}
	err := service.NewMenuService().BatchDeleteMenuByIds(req)
	if err != nil {
		t.Fatal("批量删除菜单失败: ", err)
	}
	fmt.Println("成功")
}

// 根据用户ID获取用户的权限(可访问)菜单列表
func TestGetUserMenusByUserId(t *testing.T) {
	menu, err := service.NewMenuService().GetUserMenusByUserId(1)
	if err != nil {
		t.Fatal("根据用户ID获取用户的权限(可访问)菜单列表失败: ", err)
	}
	for _, v := range menu {
		fmt.Println(*v)
	}
}

// 根据用户ID获取用户的权限(可访问)菜单树
func TestGetUserMenuTreeByUserId(t *testing.T) {
	menu, err := service.NewMenuService().GetUserMenuTreeByUserId(1)
	if err != nil {
		t.Fatal("根据用户ID获取用户的权限(可访问)菜单树失败: ", err)
	}
	for _, v := range menu {
		fmt.Println(*v)
	}
}
