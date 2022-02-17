package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// Response 基础序列化器
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

// PageResult
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
}

func Result(resp Response, c *gin.Context) {
	c.JSON(http.StatusOK, resp)
}

func respResult(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	respResult(200, map[string]interface{}{}, "ok", c)
}

func OkWithMessage(message string, c *gin.Context) {
	respResult(200, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	respResult(200, data, "ok", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	respResult(200, data, message, c)
}

func Err(c *gin.Context) {
	respResult(-1, map[string]interface{}{}, "internal server error!", c)
}

func ErrWithMessage(msg string, c *gin.Context) {
	respResult(-1, "", msg, c)
}

func ErrWithDetailed(code int, data interface{}, msg string, c *gin.Context) {
	respResult(code, data, msg, c)
}

// ParamErr 各种参数错误
func ParamErr(code int, msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	res := Response{
		Code: code,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Msg = err.Error()
	}
	return res
}

// ErrorResponse 返回错误消息
func ErrorResponse(code int, msg string, err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := T(fmt.Sprintf("Field.%s", e.Field))
			tag := T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return ParamErr(code,
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamErr(-1, "JSON类型不匹配", err)
	}
	return ParamErr(-1, "参数错误", err)
}
