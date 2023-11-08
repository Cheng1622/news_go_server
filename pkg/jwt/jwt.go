package jwt

import (
	"errors"

	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/golang-jwt/jwt"

	"time"
)

type JwtClaims struct {
	UserId int64
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userid int64) (string, error) {

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return "", err
	}

	c := JwtClaims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Conf.Jwt.Timeout) * time.Hour).Unix(), //过期时间
			Issuer:    config.Conf.Jwt.Issuer,                                                    // 签发人
			Subject:   config.Conf.Jwt.Subject,                                                   // 主题
			IssuedAt:  time.Now().Unix(),                                                         // 签发时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS512, c)
	return token.SignedString(signKey)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*JwtClaims, error) {
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(config.Conf.System.RSAPublicBytes)
	if err != nil {
		return nil, err
	}

	var mc = new(JwtClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")

}
