package helper

import (
	"net/http"
	"zliway/global"
)

/**
 * 权限相关工具
 * @author eyesYeager
 * @date 2023/4/12 11:32
 */

// CheckAuth 检查权限
func CheckAuth(r *http.Request) bool {
	return r.Header.Get(global.Config.App.TokenHeader) == global.Config.App.Token
}
