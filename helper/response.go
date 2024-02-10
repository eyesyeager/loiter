package helper

import (
	"loiter/config"
	"loiter/constant"
	"net/http"
)

/**
 * 响应体工具
 * @auth eyesYeager
 * @date 2024/1/9 16:38
 */

// Response 响应工具
func Response(w http.ResponseWriter, statusCode int, contentType string, content string) {
	w.Header().Set("content-type", contentType)
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(content))
}

// HtmlSimpleTemplate HTML简单响应模版
func HtmlSimpleTemplate(titleStruct constant.ResponseTitleStruct, msg string) (int, string, string) {
	systemName := config.Program.Name
	// 线上环境不展示具体错误信息
	if config.Program.Mode == constant.ONLINE {
		msg = constant.ResponseNotice.Empty
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
