package interceptor

import (
	"net/http"
)

/**
 * 响应拦截器
 * @author eyesYeager
 * @date 2023/11/25 21:34
 */

func ResponseInterceptor(w http.ResponseWriter, r *http.Request) {
	printRespLog(r)
}

// printRespLog 打印响应日志
func printRespLog(r *http.Request) {
	//responseDump, err := httputil.DumpResponse(r.Response, true)
	//if err != nil {
	//	global.AppLogger.Error(fmt.Sprintf("response log printing error: %s", err.Error()))
	//	return
	//}
	//global.AppLogger.Info(string(responseDump))
}
