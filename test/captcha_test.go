package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/captcha"
	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/redis"
)

type CaptchaReq struct {
	ImgHeight int `json:"imgHeight"`
	ImgWidth  int `json:"imgWidth"`
	KeyLong   int `json:"keyLong"`
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
	Captcha   string `json:"captcha"`
}

func TestCaptcha(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	redis.InitRedis()
	CaptchaReq := &captcha.CaptchaReq{
		ImgHeight: 80,
		ImgWidth:  270,
		KeyLong:   6,
	}
	data, err := captcha.GenCaptcha(CaptchaReq)
	if err != nil {
		t.Fatal("生成验证码错误:", err)
	}
	t.Fatal(data)
}
