package test

import (
	"fmt"
	"testing"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
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

// 获取角色列表
func TestGetRole(t *testing.T) {
	req := &requ.RoleListRequest{
		Name: "超级管理员",
	}
	roles, total, err := service.NewRoleService().GetRoles(req)
	if err != nil {
		t.Fatal("获取角色列表失败: ", err)
	}
	t.Fatal(roles, total)
}

// 根据角色ID获取角色
func TestGetRolesByIds(t *testing.T) {
	req := []uint{1, 2, 3, 4, 5, 6}
	roles, err := service.NewRoleService().GetRolesByIds(req)
	if err != nil {
		t.Fatal("获取角色列表失败: ", err)
	}
	for _, v := range roles {
		fmt.Println(*v)
	}
}

// 创建角色
func TestCreateRole(t *testing.T) {
	req := &model.Role{
		Name:    "VIP用户",
		Keyword: "vip_user",
		Desc:    "VIP用户",
		Status:  1,
		Sort:    4,
		Creator: "admin",
	}
	err := service.NewRoleService().CreateRole(req)
	if err != nil {
		t.Fatal("创建角色失败: ", err)
	}
	fmt.Println("创建角色成功")
}

// 更新角色
func TestUpdateRoleById(t *testing.T) {
	req := &model.Role{
		Name:    "VIP用户",
		Keyword: "vip_user",
		Desc:    "VIP用户",
		Status:  1,
		Sort:    4,
		Creator: "admin",
	}
	err := service.NewRoleService().UpdateRoleById(6, req)
	if err != nil {
		t.Fatal("根据id修改角色失败: ", err)
	}
	fmt.Println("根据id修改角色成功")
}

// 获取角色的权限菜单
func TestGetRoleMenusById(t *testing.T) {
	menu, err := service.NewRoleService().GetRoleMenusById(1)
	if err != nil {
		t.Fatal("获取角色的权限菜单失败: ", err)
	}
	for _, v := range menu {
		fmt.Println(*v)
	}
}

// 更新角色的权限菜单新角色的权限菜单
func TestUpdateRoleMenus(t *testing.T) {
	req := &model.Role{
		Name: "VIP用户",
	}
	err := service.NewRoleService().UpdateRoleMenus(req)
	if err != nil {
		t.Fatal("更新角色的权限菜单失败: ", err)
	}
	fmt.Println("成功")
}

// 根据角色关键字获取角色的权限接口
func TestGetRoleApisByRoleKeyword(t *testing.T) {
	api, err := service.NewRoleService().GetRoleApisByRoleKeyword("vip_user")
	if err != nil {
		t.Fatal("根据角色关键字获取角色的权限接口失败: ", err)
	}
	for _, v := range api {
		fmt.Println(*v)
	}
}

// 更新角色的权限接口（先全部删除再新增）
func TestUpdateRoleApis(t *testing.T) {

}

// 删除角色
func TestBatchDeleteRoleByIds(t *testing.T) {
	req := []uint{6}
	err := service.NewRoleService().BatchDeleteRoleByIds(req)
	if err != nil {
		t.Fatal("删除角色失败: ", err)
	}
	fmt.Println("成功")
}
