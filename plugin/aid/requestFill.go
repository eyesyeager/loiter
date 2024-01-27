package aid

import (
	"loiter/kernel/store"
	"net/http"
)

/**
 * 给请求日志插入耗时
 * @auth eyesYeager
 * @date 2024/1/25 16:31
 */

// RequestFillAid 请求日志填充
func RequestFillAid(w http.ResponseWriter, req *http.Request, resp *http.Response, host string) error {
	go func() {
		println(store.GetValue(req, store.RequestLogTime))
		println(store.GetValue(req, store.RequestLogId))
	}()
	return nil
}
