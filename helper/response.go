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
	return titleStruct.Status, "text/html;charset=utf-8", `
<!DOCTYPE html>
<html>
<head><title>` + titleStruct.Title + `</title></head>
<body bgcolor="white">
    <center><h1>` + titleStruct.Title + `</h1></center>
    <hr />
    <center><p>` + systemName + `</p></center>
	<br />
    <p>` + msg + `</p>
</body>
</html>
`
}
