package proxy

import (
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 初始化配置
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 进入前置管道

		// 进入过滤器

		// 进入后置管道

		// 获取代理信息
		//err, targetServer := balancer.SpecifyURL(r)
		//if err != nil {
		//	_, _ = w.Write([]byte(err.Error()))
		//	return
		//}

		// 创建代理
		//proxy := httputil.NewSingleHostReverseProxy(targetServer)

		// 进入响应管道（异步执行）

		// 执行反向代理
		//proxy.ServeHTTP(w, r)
	})
}
