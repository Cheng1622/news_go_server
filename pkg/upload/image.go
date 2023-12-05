package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/encrypt"
	"github.com/Cheng1622/news_go_server/pkg/file"
)

// GetImageFullUrl 获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return config.Conf.Upload.ImagePrefixUrl + "/" + GetImagePath() + name
}

// GetImageName 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = encrypt.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath 获取图片路径
func GetImagePath() string {
	return config.Conf.Upload.ImageSavePath
}

// CheckImageExt 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range config.Conf.Upload.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		clog.Log.Errorln(err)
		return false
	}

	return size <= config.Conf.Upload.ImageMaxSize*1024*1024
}

// CheckImage 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
