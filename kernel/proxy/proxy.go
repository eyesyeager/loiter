package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"zliway/global"
)

/**
 * 实现代理
 * @author eyesYeager
 * @date 2023/4/9 21:08
 */

// CreateProxy 创建反向代理
func CreateProxy() *http.ServeMux {
	mux := http.NewServeMux()

	// 注册不同的服务路由
	mux.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		proxy := newSingleHostReverseProxy("http://localhost:9501")
		proxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		proxy := newSingleHostReverseProxy("http://localhost:9502")
		proxy.ServeHTTP(w, r)
	})

	return mux
}

// 创建代理

// 代理创建工具方法
func newSingleHostReverseProxy(target string) *httputil.ReverseProxy {
	pUrl, err := url.Parse(target)
	if err != nil {
		global.Log.Error("parses a raw url '" + target + "' into a URL structure failed: " + err.Error())
	}
	return httputil.NewSingleHostReverseProxy(pUrl)
}
