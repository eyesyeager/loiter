package result

import (
	"encoding/json"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/4/11 21:50
 */

// ResponseUtil 接口返回值封装(默认StatusOK状态码)
func ResponseUtil(w http.ResponseWriter, r Response) {
	ResponseUtilWithStatus(w, r, http.StatusOK)
}

// ResponseUtilWithStatus 接口返回值封装(自定义状态码)
func ResponseUtilWithStatus(w http.ResponseWriter, r Response, code int) {
	w.Header().Set("content-type", "text/json")
	msg, _ := json.Marshal(r)
	w.WriteHeader(code)
	_, _ = w.Write(msg)
}
