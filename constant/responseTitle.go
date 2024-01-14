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
}

type responseTitle struct {
	BadGateway ResponseTitleStruct
}

type ResponseTitleStruct struct {
	Title  string
	Status int
}
