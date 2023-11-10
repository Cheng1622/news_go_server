package captcha

import (
	"github.com/mojocn/base64Captcha"
)

// CaptchaReq
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

// 生成验证码
func GenCaptcha(ca *CaptchaReq) (result CaptchaResponse, err error) {
	// 图片高度80 宽度240 数字位数6 最大绝对偏斜因子 背景圆圈数量
	driver := base64Captcha.NewDriverDigit(ca.ImgHeight, ca.ImgWidth, ca.KeyLong, 0.7, 80)
	//base64Captcha.NewDriverMath()
	// 生成验证码并保存至store
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	// 生成base64图像及id
	id, b64s, err := cp.Generate()
	if err != nil {
		return result, err
	}
	captcha := base64Captcha.DefaultMemStore.Get(id, true)

	// 数字验证码存redis
	CaptchaCache := NewCaptchaService()
	_ = CaptchaCache.SetCaptcha(id, captcha)
	// 返回值
	result = CaptchaResponse{
		CaptchaId: id,
		PicPath:   b64s,
		Captcha:   captcha,
	}

	return result, nil
}

// VerifyCaptcha 验证验证码 原验证码的id，待验证的输入字符串answer
func VerifyCaptcha(id string, answer string) bool {
	return base64Captcha.DefaultMemStore.Verify(id, answer, true)
}
