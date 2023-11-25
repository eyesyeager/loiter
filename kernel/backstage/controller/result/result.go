package result

import (
	"net/http"
)

/**
 * 返回结果封装
 * @author eyesYeager
 * @date 2023/4/11 21:22
 */

// Response 响应结构体
type Response struct {
	Code int         `json:"code"` // 自定义错误码
	Msg  string      `json:"msg"`  // 信息
	Data interface{} `json:"data"` // 数据
}

// Success 成功响应
func Success(w http.ResponseWriter, r *http.Request, code int, msg string, data interface{}) {
	result := Response{
		code,
		msg,
		data,
	}
	ResponseUtil(w, result)
}

// SuccessByCustom 成功响应(使用customResult信息)
func SuccessByCustom(w http.ResponseWriter, r *http.Request, result customResult, data interface{}) {
	Success(w, r, result.Code, result.Msg, data)
}

// SuccessDefault 成功响应(默认模式)
func SuccessDefault(w http.ResponseWriter, r *http.Request, data interface{}) {
	SuccessByCustom(w, r, Results.DefaultSuccess, data)
}

// SuccessAttachedCode 成功响应(默认模式，自选状态码)
func SuccessAttachedCode(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	Success(w, r, code, Results.DefaultSuccess.Msg, data)
}

// SuccessAttachedMsg 成功响应(默认模式，自选信息)
func SuccessAttachedMsg(w http.ResponseWriter, r *http.Request, data interface{}, msg string) {
	Success(w, r, Results.DefaultSuccess.Code, msg, data)
}

// Fail 失败响应
func Fail(w http.ResponseWriter, r *http.Request, code int, msg string) {
	result := Response{
		code,
		msg,
		nil,
	}
	ResponseUtil(w, result)
}

// FailByCustom 失败响应(使用customResult信息)
func FailByCustom(w http.ResponseWriter, r *http.Request, result customResult) {
	Fail(w, r, result.Code, result.Msg)
}

// FailDefault 失败响应(默认模式)
func FailDefault(w http.ResponseWriter, r *http.Request) {
	FailByCustom(w, r, Results.DefaultFail)
}

// FailAttachedCode 失败响应(默认模式，自选状态码)
func FailAttachedCode(w http.ResponseWriter, r *http.Request, code int) {
	Fail(w, r, code, Results.DefaultFail.Msg)
}

// FailAttachedMsg 失败响应(默认模式，自选信息)
func FailAttachedMsg(w http.ResponseWriter, r *http.Request, msg string) {
	Fail(w, r, Results.DefaultFail.Code, msg)
}
