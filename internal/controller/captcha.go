package controller

import (
	"github.com/Cheng1622/news_go_server/pkg/captcha"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/gin-gonic/gin"
)

type CaptchaService interface {
	Captcha(c *gin.Context) // 获取接口列表
}

// CaptchaApiService 服务层数据处理
type CaptchaApiService struct{}

// NewCaptchaApi 创建构造函数简单工厂模式
func NewCaptchaApi() CaptchaService {
	return CaptchaApiService{}
}

// Captcha 生成验证码
// @Tags		Base
// @Summary	生成验证码
// @accept		application/json
// @Produce	application/json
// @Success	1000	{object}	response.Response{data=captcha.CaptchaResponse}	"生成验证码,返回包括随机数id,base64,验证码长度"
// @Router		/api/v1/base/captcha [get]
func (cs CaptchaApiService) Captcha(c *gin.Context) {
	CaptchaReq := &captcha.CaptchaReq{
		ImgHeight: 80,
		ImgWidth:  270,
		KeyLong:   6,
	}
	data, err := captcha.GenCaptcha(CaptchaReq)
	if err != nil {
		clog.Log.Errorln("生成验证码错误:", err)
		response.Error(c, code.ServerErr, nil)
		return
	}
	response.Success(c, code.SUCCESS, data)
}
