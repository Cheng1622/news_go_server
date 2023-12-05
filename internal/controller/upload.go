package controller

import (
	"github.com/Cheng1622/news_go_server/internal/service"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/gin-gonic/gin"
)

type UploadService interface {
	UploadImage(c *gin.Context) // 上传图片
}

// UploadApiService 服务层数据处理
type UploadApiService struct {
	Upload service.UploadService
}

// NewUploadApi 创建构造函数简单工厂模式
func NewUploadApi() UploadService {
	return UploadApiService{Upload: service.NewUploadService()}
}

// 上传图片
func (u UploadApiService) UploadImage(c *gin.Context) {

	file, err := u.Upload.UploadImage(c)

	if err != nil {
		response.Error(c, code.ServerErr, err.Error())
		return
	}

	// 成功返回
	response.Success(c, code.SUCCESS, file)
	return

}
