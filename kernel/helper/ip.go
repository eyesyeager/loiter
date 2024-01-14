package helper

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

/**
 * ip工具
 * @author eyesYeager
 * @date 2023/4/10 9:42
 */

// GetIp 获取用户ip
func GetIp(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
