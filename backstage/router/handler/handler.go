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
	interceptor.RequestInterceptor(w, req)
	r.router.ServeHTTP(w, req)
	interceptor.ResponseInterceptor(w, req)
}

func New(router *httprouter.Router) *LoiterRouter {
	return &LoiterRouter{
		router: router,
	}
}
