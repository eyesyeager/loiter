package router

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/router/route"
)

/**
 * 后台管理路由
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

// InitRouter 初始化路由
func InitRouter(routerRoot *httprouter.Router) {
	route.InitUserRoute(routerRoot, "/user")
	route.InitLogRoute(routerRoot, "/log")
	route.InitAppRoute(routerRoot, "/app")
	route.InitServerRoute(routerRoot, "/server")
	route.InitBalancerRoute(routerRoot, "/balancer")
	route.InitContainerRoute(routerRoot, "/container")
	route.InitPassagewayRoute(routerRoot, "/passageway")
	route.InitAidRoute(routerRoot, "/aid")
}
