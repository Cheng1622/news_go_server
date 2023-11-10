package captcha

import (
	"context"
	"fmt"
	"time"

	"github.com/Cheng1622/news_go_server/pkg/redis"
)

// CaptchaService 数字验证码缓存模块接口
type CaptchaService interface {
	SetCaptcha(id string, answer string) error // 数字验证码存 set key
	GetCaptcha(id string) string               // 数字验证码存 get key

}

type Captcha struct{}

// NewCaptchaService 构造函数
func NewCaptchaService() CaptchaService {
	return Captcha{}
}

// SetCaptcha  数字验证码存 set key
func (ca Captcha) SetCaptcha(id string, answer string) error {
	// 数字验证码 key
	key := fmt.Sprintf("captcha:%s", id)
	err := redis.Redis.Set(context.Background(), key, answer, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetCaptcha  数字验证码 get key
func (ca Captcha) GetCaptcha(id string) string {
	// 数字验证码 key
	key := fmt.Sprintf("captcha:%s", id)
	captcha, _ := redis.Redis.Get(context.Background(), key).Result()
	return captcha
}
