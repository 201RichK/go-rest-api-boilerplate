package controllers

import (
	"net/http"

	"github.com/aro/infra/logger"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

// Response object as HTTP response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// HTTPRes normalize HTTP Response format
func HTTPRes(ctx *gin.Context, httpCode int, msg string, data interface{}) {
	if httpCode == http.StatusInternalServerError {
		logger.Error(msg)
		msg = ginI18n.MustGetMessage(ctx, "internal_server_error")
	}

	ctx.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})
}
