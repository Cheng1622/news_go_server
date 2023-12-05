package test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/validator"
	"github.com/gin-gonic/gin"
)

func TestValidator(t *testing.T) {
	config.InitConfig()
	clog.InitLogger()
	validator.InitValidate()

	// 创建一个 gin 的上下文
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// 模拟一个验证错误
	err := errors.New("Invalid format")
	validator.HandleValidatorError(c, err)
	body := w.Body.String()
	// 打印响应数据
	t.Fatal("Response Body:", body)
}
