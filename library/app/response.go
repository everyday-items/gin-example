package app

import (
	"github.com/everyday-items/gin-example/library/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

// ResponseWithMsg 自定义消息响应
func (g *Gin) ResponseWithMsg(httpCode, errCode int, msg string, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}
