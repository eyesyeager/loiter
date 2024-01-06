package proxy

import (
	"loiter/kernel/balancer"
	"loiter/kernel/pipeline"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 初始化配置
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host := r.Host

		// 进入前置管道
		pipeline.FrontEntry(host, w, r)

		// 进入过滤器

		// 进入后置管道
		pipeline.RearEntry(host, w, r)

		// 获取代理信息
		targetUrl := balancer.Entry(host)
		println("-----------")
		println(targetUrl)
		println("-----------")

		// 创建代理
		//proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		// 进入响应管道（异步执行）

		// 执行反向代理
		//proxy.ServeHTTP(w, r)
	})
}
