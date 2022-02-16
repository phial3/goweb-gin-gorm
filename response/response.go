package response

import (
	"encoding/json"
	"fmt"
)

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

import (
	"goweb-gin-gorm/constant"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: constant.CodeCheckLogin,
		Msg:  "未登录",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(constant.CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(constant.CodeParamErr, msg, err)
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := T(fmt.Sprintf("Field.%s", e.Field))
			tag := T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamErr("JSON类型不匹配", err)
	}
	return ParamErr("参数错误", err)
}
