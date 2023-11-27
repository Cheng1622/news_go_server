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

// 获取接口列表
func TestGetApis(t *testing.T) {
	req := &requ.ApiListRequest{
		Method:   "",
		Path:     "",
		Category: "",
		Creator:  "",
		PageNum:  1,
		PageSize: 2,
	}
	api, total, err := service.NewApiService().GetApis(req)
	if err != nil {
		t.Fatal("获取接口列表失败: ", err)
	}
	for _, v := range api {
		fmt.Println(*v)
	}
	fmt.Println("total:", total)
}

// 根据接口ID获取接口列表
func TestGetApisById(t *testing.T) {
	req := []uint{1}
	api, err := service.NewApiService().GetApisById(req)
	if err != nil {
		t.Fatal("根据接口ID获取接口列表失败: ", err)
	}
	for _, v := range api {
		fmt.Println(*v)
	}
}

// 获取接口树
func TestGetApiTree(t *testing.T) {
	api, err := service.NewApiService().GetApiTree()
	if err != nil {
		t.Fatal("获取接口树失败: ", err)
	}
	for _, v := range api {
		fmt.Println(*v)
	}
}

// 创建接口
func TestCreateApi(t *testing.T) {
	req := &model.Api{
		Method:   "",
		Path:     "",
		Category: "",
		Desc:     "",
		Creator:  "",
	}
	err := service.NewApiService().CreateApi(req)
	if err != nil {
		t.Fatal("创建接口失败: ", err)
	}
	fmt.Println("成功")
}

// 更新接口
func TestUpdateApiById(t *testing.T) {
	req := &model.Api{
		Method:   "POST",
		Path:     "/base/login",
		Category: "base",
		Desc:     "用户登录",
		Creator:  "admin",
	}
	err := service.NewApiService().UpdateApiById(1, req)
	if err != nil {
		t.Fatal("更新接口失败: ", err)
	}
	fmt.Println("成功")
}

// 批量删除接口
func TestBatchDeleteApiByIds(t *testing.T) {
	req := []uint{1}
	err := service.NewApiService().BatchDeleteApiByIds(req)
	if err != nil {
		t.Fatal("批量删除接口失败: ", err)
	}
	fmt.Println("成功")
}

// 根据接口路径和请求方式获取接口描述
func TestGetApiDescByPath(t *testing.T) {
	method := "POST"
	path := "/base/login"
	api, err := service.NewApiService().GetApiDescByPath(path, method)
	if err != nil {
		t.Fatal("根据接口路径和请求方式获取接口描述失败: ", err)
	}
	fmt.Println(api)
}
