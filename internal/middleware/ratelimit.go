package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimitMiddleware fillInterval 指每过多长时间向桶里放一个令牌，capacity 是桶的容量
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌 就直接返回吧
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "访问限流")
			c.Abort()
			return
		}
		c.Next()
	}
}
