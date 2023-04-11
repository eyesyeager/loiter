package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

/**
 * 实现代理
 * @author eyesYeager
 * @date 2023/4/9 21:08
 */

func main() {
	mux := http.NewServeMux()

	// 注册不同的服务路由
	mux.HandleFunc("service1.com/", func(w http.ResponseWriter, r *http.Request) {
		proxy := NewSingleHostReverseProxy("http://localhost:8080")
		proxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("service2.com/", func(w http.ResponseWriter, r *http.Request) {
		proxy := NewSingleHostReverseProxy("http://localhost:8081")
		proxy.ServeHTTP(w, r)
	})

	_ = http.ListenAndServe(":80", mux)
}

// NewSingleHostReverseProxy 创建反向代理
func NewSingleHostReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}
