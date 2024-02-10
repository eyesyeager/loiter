package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 过滤器路由
 * @auth eyesYeager
 * @date 2024/1/11 18:53
 */

func InitFilterRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/updateAppFilter", controller.UpdateAppFilter)
	routerRoot.POST(group+"/updateAppLimiter", controller.UpdateAppLimiter)
	routerRoot.POST(group+"/updateAppNameList", controller.UpdateAppNameList)
	routerRoot.POST(group+"/addNameListIp", controller.AddNameListIp)
	routerRoot.POST(group+"/deleteNameListIp", controller.DeleteNameListIp)
}
