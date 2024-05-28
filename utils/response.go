package utils

import (
	"encoding/json"
	"loiter/config"
	"loiter/constants"
	"net/http"
)

/**
 * 响应体工具
 * @auth eyesYeager
 * @date 2024/1/9 16:38
 */

// responseJsonStruct json格式响应结构体
type responseJsonStruct struct {
	Code int    `json:"code"` // 自定义错误码
	Msg  string `json:"msg"`  // 信息
}

// Response 响应工具
func Response(w http.ResponseWriter, statusCode int, contentType string, content string) {
	w.Header().Set("content-type", contentType)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
	w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(content))
}

// ResponseTemplate 响应模版
func ResponseTemplate(titleStruct constants.ResponseTitleStruct, msg string, genre string) (int, string, string) {
	if genre == constants.AppGenre.Api {
		return JsonTemplate(titleStruct, msg)
	}
	return HtmlSimpleTemplate(titleStruct, msg)
}

// HtmlSimpleTemplate HTML简单响应模版
func HtmlSimpleTemplate(titleStruct constants.ResponseTitleStruct, msg string) (int, string, string) {
	systemName := config.Program.Name
	// 线上环境不展示具体错误信息
	if config.Program.Mode == constants.ONLINE {
		msg = constants.ResponseNotice.Empty
	}
	return titleStruct.Status, "text/html;charset=utf-8", `
<html>
<head><title>` + titleStruct.Title + `</title></head>
<body>
    <center><h1>` + titleStruct.Title + `</h1></center>
    <hr />
    <center>` + systemName + `</center>
	<br />
    <p>` + msg + `</p>
</body>
</html>
`
}

// JsonTemplate json响应模版
func JsonTemplate(titleStruct constants.ResponseTitleStruct, msg string) (int, string, string) {
	// 线上环境不展示具体错误信息
	if config.Program.Mode == constants.ONLINE {
		msg = titleStruct.Title
	}
	marshal, _ := json.Marshal(responseJsonStruct{
		Code: titleStruct.Status,
		Msg:  msg,
	})
	return titleStruct.Status, "text/json", string(marshal)
}
