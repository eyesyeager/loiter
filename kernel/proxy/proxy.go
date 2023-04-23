package proxy

import (
	"net/http"
	"net/http/httputil"
	"zliway/kernel/balancer"
)

/**
 * 实现代理
 * @author eyesYeager
 * @date 2023/4/9 21:08
 */

// StartProxy 启动反向代理
func StartProxy() {
	// 初始化配置
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 进入过滤器

		// 获取代理信息
		err, targetServer := balancer.SpecifyURL(r)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		// 创建代理
		proxy := httputil.NewSingleHostReverseProxy(targetServer)

		// 执行反向代理
		proxy.ServeHTTP(w, r)
	})
}
