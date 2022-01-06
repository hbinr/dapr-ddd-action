package httpx

import (
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func RespSuccess(data interface{}) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
	}
}
