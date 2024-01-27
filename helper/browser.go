package helper

import "net/http"

/**
 * 浏览器工具
 * @author eyesYeager
 * @date 2023/4/13 11:13
 */

// GetBrowser 获取浏览器信息
func GetBrowser(r *http.Request) string {
	return r.Header.Get("User-Agent")
}
