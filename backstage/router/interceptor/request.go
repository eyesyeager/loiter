package interceptor

import (
	"fmt"
	"loiter/global"
	"net/http"
	"net/http/httputil"
)

/**
 * 请求拦截器
 * @author eyesYeager
 * @date 2023/11/25 21:34
 */

func RequestInterceptor(w http.ResponseWriter, r *http.Request) bool {
	if !cors(w, r) {
		return false
	}
	printReqLog(r)
	return true
}

// printReqLog 打印请求日志
func printReqLog(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		global.AppLogger.Error(fmt.Sprintf("request log printing error: %s", err.Error()))
		return
	}
	global.AppLogger.Info(string(requestDump))
}

// cors 防跨域
func cors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 放行所有 OPTIONS 方法
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return false
	}
	return true
}
