package controllers

import (
	"go_server/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListHandler(c *gin.Context) {
	data, err := logic.GetListList()
	if err != nil {
		zap.L().Error("GetListList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func ListLastHandler(c *gin.Context) {
	data, err := logic.GetListListLast()
	if err != nil {
		zap.L().Error("GetListList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func ListDetailHandler(c *gin.Context) {
	listIdStr := c.Param("id")
	listId, err := strconv.ParseInt(listIdStr, 10, 64)
	// 校验参数是否正确
	if err != nil {
		zap.L().Error("GetListListDetail error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetListById(listId)
	if err != nil {
		zap.L().Error("GetListListDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
