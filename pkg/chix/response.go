package chix

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/dapr-ddd-action/app/pkg/constant/e"
)

// Response .
type Response struct {
	Code e.ResCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// RespSuccess 响应成功
func RespSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, &Response{
		Code: e.CodeSuccess,
		Msg:  e.CodeSuccess.Msg(),
		Data: data,
	})
	return
}

// RespError 响应失败，携带状态及对应信息
func RespError(w http.ResponseWriter, r *http.Request, code e.ResCode) {
	render.JSON(w, r, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
	return
}

// RespErrorWithMsg 响应失败，携带状态+其他自定义信息
func RespErrorWithMsg(w http.ResponseWriter, r *http.Request, code e.ResCode, msg string) {
	render.JSON(w, r, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
	return
}
