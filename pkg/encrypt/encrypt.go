package encrypt

import "github.com/Cheng1622/news_go_server/pkg/config"

// 密码加密
func NewGenPasswd(passwd string) string {
	pass := RSAEncrypt([]byte(passwd), config.Conf.System.RSAPublicBytes)
	return string(pass)
}

// 密码解密
func NewParPasswd(passwd string) string {
	pass := RSADecrypt([]byte(passwd), config.Conf.System.RSAPrivateBytes)
	return string(pass)
}
