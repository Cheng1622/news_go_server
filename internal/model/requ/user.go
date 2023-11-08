package requ

// 用户登录结构体
type RegisterAndLoginRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`
}
