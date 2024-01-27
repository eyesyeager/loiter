package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/controller"
)

/**
 * 通道路由
 * @auth eyesYeager
 * @date 2024/1/11 18:53
 */

func InitPassagewayRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/updateAppPassageway", controller.UpdateAppPassageway)
	routerRoot.POST(group+"/updateAppLimiter", controller.UpdateAppLimiter)
	routerRoot.POST(group+"/updateAppNameList", controller.UpdateAppNameList)
	routerRoot.POST(group+"/addNameListIp", controller.AddNameListIp)
	routerRoot.POST(group+"/deleteNameListIp", controller.DeleteNameListIp)
}
