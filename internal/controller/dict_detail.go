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

// DictDetailsService
type DictDetailsService interface {
	PostDictDetails(c *gin.Context)      // 创建
	GetDictDetailsList(c *gin.Context)   // 列表
	PutDictDetails(c *gin.Context)       // 更新
	DeleteDictDetails(c *gin.Context)    // 删除
	DeleteDictDetailsAll(c *gin.Context) // 批量删除
}

// DictDetailsApiService 服务层数据处理
type DictDetailsApiService struct {
	DictDetails service.DictDetailsService
}

// NewDictDetailsApi 构造函数
func NewDictDetailsApi() DictDetailsService {
	return DictDetailsApiService{DictDetails: service.NewDictDetailsService()}
}

// PostDictDetails 创建
func (ds DictDetailsApiService) PostDictDetails(c *gin.Context) {
	example := new(model.DictDetail)
	// 参数绑定
	if err := c.ShouldBindJSON(&example); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.DictDetails.PostDictDetails(example)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// GetDictDetailsList 列表
func (ds DictDetailsApiService) GetDictDetailsList(c *gin.Context) {

	var pageList requ.DictDetailList
	// 参数绑定
	if err := c.ShouldBindQuery(&pageList); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 参数校验

	// 服务层数据操作
	data, total, err := ds.DictDetails.GetDictDetailsList(&pageList)

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

// PutDictDetails 更新
func (ds DictDetailsApiService) PutDictDetails(c *gin.Context) {
	example := new(model.DictDetail)
	// 参数绑定
	if err := c.ShouldBind(&example); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.DictDetails.PutDictDetails(example)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// DeleteDictDetails 删除
func (ds DictDetailsApiService) DeleteDictDetails(c *gin.Context) {
	DictId := new(requ.DictId)
	if err := c.ShouldBind(&DictId); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.DictDetails.DeleteDictDetails(DictId)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}

// DeleteDictDetailsAll 批量删除
func (ds DictDetailsApiService) DeleteDictDetailsAll(c *gin.Context) {
	dictDetaiIds := new(requ.DictDetaiIds)
	// 参数绑定
	if err := c.ShouldBind(&dictDetaiIds); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}
	// 服务层数据操作
	err := ds.DictDetails.DeleteDictDetailsAll(dictDetaiIds.DictDetaiIds)
	if err != nil {
		// 错误返回
		response.Error(c, code.ServerErr, err.Error())
		return
	}
	// 成功返回
	response.Success(c, code.SUCCESS, nil)
	return
}
