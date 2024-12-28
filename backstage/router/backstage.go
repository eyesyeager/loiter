package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"loiter/backstage/router/route"
	"loiter/config"
	"net/http"
)

/**
 * 后台管理路由
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

// InitRouter 初始化路由
func InitRouter(routerRoot *httprouter.Router) {
	// 后台程序
	route.InitUserRoute(routerRoot, "/user")
	route.InitLogRoute(routerRoot, "/log")
	route.InitAppRoute(routerRoot, "/app")
	route.InitBalancerRoute(routerRoot, "/balancer")
	route.InitContainerRoute(routerRoot, "/container")
	route.InitProcessorRoute(routerRoot, "/processor")
	route.InitCommonRoute(routerRoot, "/common")
	route.InitNoticeRoute(routerRoot, "/notice")

	// 监控数据
	if config.Program.PrometheusConfig.Enabled {
		routerRoot.GET(config.Program.PrometheusConfig.Path, adaptHandler(promhttp.Handler()))
	}
}

// 适配 httprouter
func adaptHandler(handler http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler.ServeHTTP(w, r)
	}
}
