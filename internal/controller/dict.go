package controller

import (
	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/Cheng1622/news_go_server/pkg/validator"
	"github.com/gin-gonic/gin"
)

// DictService
type DictService interface {
	PostDict(c *gin.Context)      // 创建
	GetDictList(c *gin.Context)   // 列表
	PutDict(c *gin.Context)       // 更新
	DeleteDict(c *gin.Context)    // 删除
	DeleteDictAll(c *gin.Context) // 批量删除
}

// DictApiService 服务层数据处理
type DictApiService struct {
	Dict service.DictService
}

// NewDictApi 构造函数
func NewDictApi() DictService {
	return DictApiService{Dict: service.NewDictService()}
}

// PostDict 创建
func (ds DictApiService) PostDict(c *gin.Context) {
	dict := new(model.Dict)
	// 参数绑定
	if err := c.ShouldBindJSON(&dict); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.Dict.PostDict(dict)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// GetExampleList 列表
func (ds DictApiService) GetDictList(c *gin.Context) {
	var pageList requ.PageList
	// 参数绑定
	if err := c.ShouldBindQuery(&pageList); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 参数校验

	// 服务层数据操作
	data, total, err := ds.Dict.GetDictList(&pageList)

	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, map[string]interface{}{
		"data":  data,
		"total": total,
	})
	return
}

// PutExample 更新
func (ds DictApiService) PutDict(c *gin.Context) {
	dict := new(model.Dict)
	// 参数绑定
	if err := c.ShouldBind(&dict); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.Dict.PutDict(dict)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// DeleteDict  删除
func (ds DictApiService) DeleteDict(c *gin.Context) {
	dictId := new(requ.DictId)
	if err := c.ShouldBind(&dictId); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.Dict.DeleteDict(dictId)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// DeleteDictAll  批量删除
func (ds DictApiService) DeleteDictAll(c *gin.Context) {
	dictIds := new(requ.DictIds)
	// 参数绑定
	if err := c.ShouldBind(&dictIds); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.Dict.DeleteDictAll(dictIds.DictIds)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, nil)
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}
