package handler

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/router/interceptor"
	"net/http"
)

/**
 * 自定义处理器
 * @auth eyesYeager
 * @date 2024/2/7 17:47
 */

type LoiterRouter struct {
	router *httprouter.Router
}

// ServeHTTP 让处理器实现 http.Handler 接口.
func (r *LoiterRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 进入请求拦截器
	if !interceptor.RequestInterceptor(w, req) {
		return
	}
	r.router.ServeHTTP(w, req)
	// 进入响应拦截器
	interceptor.ResponseInterceptor(w, req)
}

func New(router *httprouter.Router) *LoiterRouter {
	return &LoiterRouter{
		router: router,
	}
}
