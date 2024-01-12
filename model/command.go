package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorInfo struct {
}
type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrorResp(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, &ApiResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func SuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ApiResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
