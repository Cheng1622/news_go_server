package jwt

import (
	"context"
	"errors"
	"strconv"

	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"time"
)

type JwtClaims struct {
	Userid int64
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userid int64) (string, error) {

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return "", err
	}

	c := JwtClaims{
		Userid: userid,
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

	return nil, errors.New("token错误")

}

// JoinBlackList token 加入黑名单
func JoinBlackList(c *gin.Context) (err error) {
	tokenStr, ok := c.Get("tokenStr")
	expiresAt, ok := c.Get("expiresAt")
	if !ok {
		return errors.New("用户未登录")
	}
	nowUnix := time.Now().Unix()
	timer := time.Duration(expiresAt.(int64)-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = redis.Redis.SetNX(context.Background(), tokenStr.(string), nowUnix, timer).Err()
	return err
}

// IsInBlacklist token 是否在黑名单中
func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := redis.Redis.Get(context.Background(), tokenStr).Result()
	if joinUnixStr == "" || err != nil {
		return false
	}
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < config.Conf.Jwt.Blacktime {
		return false
	}
	return true
}
