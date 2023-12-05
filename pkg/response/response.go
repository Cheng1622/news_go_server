package response

import (
	"net/http"

	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code code.ResCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data"`
}

// Success 返回成功
func Success(c *gin.Context, code code.ResCode, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,       // 自定义code
		Msg:  code.Msg(), // message
		Data: data,       // 数据
	})
}

// Error 返回失败
func Error(c *gin.Context, code code.ResCode, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,       // 自定义code
		Msg:  code.Msg(), // message
		Data: data,       // 数据
	})
}
