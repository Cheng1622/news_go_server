package requ

// 用户登录结构体
type RegisterAndLoginRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`
}

// 创建用户结构体
type CreateUserRequest struct {
	Userid       int64  `form:"userid" json:"userid"`
	Username     string `form:"username" json:"username" validate:"required,min=2,max=20"`
	Password     string `form:"password" json:"password"  validate:"required,min=6"`
	Mobile       string `form:"mobile" json:"mobile" validate:"required,checkMobile"`
	Avatar       string `form:"avatar" json:"avatar"`
	Nickname     string `form:"nickname" json:"nickname" validate:"min=0,max=20"`
	Introduction string `form:"introduction" json:"introduction" validate:"min=0,max=255"`
	Status       uint   `form:"status" json:"status" validate:"oneof=1 2"`
	RoleIds      []uint `form:"roleIds" json:"roleIds" validate:"required"`
}

// 获取用户列表结构体
type UserListRequest struct {
	Userid   int64  `json:"userid" form:"userid"`
	Username string `json:"username" form:"username" `
	Mobile   string `json:"mobile" form:"mobile" `
	Nickname string `json:"nickname" form:"nickname" `
	Status   uint   `json:"status" form:"status" `
	PageNum  uint   `json:"pageNum" form:"pageNum"`
	PageSize uint   `json:"pageSize" form:"pageSize"`
}

// 批量删除用户结构体
type DeleteUserRequest struct {
	Userids []uint `json:"userIds" form:"userIds"`
}

// 更新密码结构体
type ChangePwdRequest struct {
	OldPassword string `json:"oldPassword" form:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" validate:"required"`
}
