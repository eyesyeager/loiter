package router

import (
	"github.com/julienschmidt/httprouter"
	"zliway/kernel/backstage/router/route"
)

/**
 * 后台管理路由
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

// InitRouter 初始化路由
func InitRouter(routerRoot *httprouter.Router) {
	// 接口路由
	route.InitAppRoute(routerRoot, "/app")
}
