package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type JsonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJsonExit(JsonRes{Code: 0, Message: "success", Data: data})
}

func Error(r *ghttp.Request, code int, msg string) {
	r.Response.WriteJsonExit(JsonRes{Code: code, Message: msg})
}

func ErrorGateway(r *ghttp.Request, svcName string, err error) {
	g.Log().Errorf(r.Context(), "[Gateway] call %s failed: %v", svcName, err)
	r.Response.WriteJsonExit(JsonRes{Code: 502, Message: "service unavailable: " + svcName})
}
