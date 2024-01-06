package balancer

import "net/url"

/**
 * 负载均衡器入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// Entry 进入负载均衡
func Entry(host string) *url.URL {
	return &url.URL{
		Host: "",
	}
}
