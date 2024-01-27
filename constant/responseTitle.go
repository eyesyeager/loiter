package constant

import "net/http"

/**
 * 通用响应标题信息
 * @auth eyesYeager
 * @date 2024/1/9 17:04
 */

var ResponseTitle = responseTitle{
	BadGateway: ResponseTitleStruct{
		Title:  "502 Bad Gateway",
		Status: http.StatusBadGateway,
	},
	RateLimit: ResponseTitleStruct{
		Title:  "429 Too Many Requests",
		Status: http.StatusTooManyRequests,
	},
	Forbidden: ResponseTitleStruct{
		Title:  "403 Forbidden",
		Status: http.StatusForbidden,
	},
}

type responseTitle struct {
	BadGateway ResponseTitleStruct
	RateLimit  ResponseTitleStruct
	Forbidden  ResponseTitleStruct
}

type ResponseTitleStruct struct {
	Title  string
	Status int
}
