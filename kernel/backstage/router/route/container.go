package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/controller"
)

/**
 * 容器相关路由
 * @auth eyesYeager
 * @date 2024/1/8 20:01
 */

func InitContainerRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.GET(group+"/refreshAllContainer/:appId", controller.RefreshAllContainer)
	routerRoot.GET(group+"/refreshAppServer/:appId", controller.RefreshAppServer)
	routerRoot.GET(group+"/refreshBalancer/:appId", controller.RefreshBalancer)
	routerRoot.GET(group+"/refreshPassageway/:appId", controller.RefreshPassageway)
	routerRoot.GET(group+"/refreshLimiter/:appId", controller.RefreshLimiter)
	routerRoot.GET(group+"/refreshNameList/:appId", controller.RefreshNameList)
	routerRoot.GET(group+"/refreshAid/:appId", controller.RefreshAid)
}
