package service

import (
	"errors"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/upload"
	"github.com/gin-gonic/gin"
)

// UploadService
type UploadService interface {
	UploadImage(c *gin.Context) (map[string]string, error) // 上传图片
}

type Upload struct{}

// NewExampleService 构造函数
func NewUploadService() UploadService {
	return Upload{}
}

// 上传图片
func (u Upload) UploadImage(c *gin.Context) (map[string]string, error) {
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		clog.Log.Errorln("上传图片错误: ", err)
		return nil, err
	}

	imageName := upload.GetImageName(image.Filename)
	savePath := upload.GetImagePath()
	fullPath := upload.GetImagePath()
	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		return nil, errors.New("图片格式或大小有问题")
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		clog.Log.Errorln("检查图片失败", err)
		return nil, errors.New("检查图片失败")
	}
	err = c.SaveUploadedFile(image, src)
	if err != nil {
		clog.Log.Errorln("检查图片失败", err)
		return nil, errors.New("检查图片失败")
	}

	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	return data, nil
}
