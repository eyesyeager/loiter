package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 负载均衡模块路由
 * @auth eyesYeager
 * @date 2024/1/5 16:50
 */

func InitBalancerRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/updateAppBalancer", controller.UpdateAppBalancer)
	routerRoot.POST(group+"/getBalancerByPage", controller.GetBalancerByPage)
}
