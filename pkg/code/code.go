package code

type ResCode int64

const (
	SUCCESS ResCode = 1000 + iota
	ERROR
	ServerErr
	ValidateError
	Deadline
	CreateError
	FindError
	WithoutServer
	AuthError
	DeleteError
	EmptyFile
	RateLimit
	Unauthorized
	WithoutLogin
	DisableAuth
	ServerBusy
)

var codeMsgMap = map[ResCode]string{
	SUCCESS:       "成功",
	ERROR:         "失败",
	ServerErr:     "服务器错误",
	ValidateError: "参数校验错误",
	Deadline:      "服务调用超时",
	CreateError:   "服务器写入失败",
	FindError:     "服务器查询失败",
	WithoutServer: "服务未启用",
	AuthError:     "权限错误",
	DeleteError:   "服务器删除失败",
	EmptyFile:     "文件为空",
	RateLimit:     "访问限流",
	Unauthorized:  "JWT认证失败",
	WithoutLogin:  "用户未登录",
	DisableAuth:   "当前用户已被禁用",
	ServerBusy:    "服务器繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[ServerBusy]
	}
	return msg
}
