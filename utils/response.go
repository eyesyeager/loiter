package utils

import (
	"loiter/config"
	"loiter/constants"
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
