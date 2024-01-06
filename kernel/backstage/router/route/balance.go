package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/controller"
)

/**
 * 负载均衡模块路由
 * @auth eyesYeager
 * @date 2024/1/5 16:50
 */

func InitBalanceRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/updateAppBalance", controller.UpdateAppBalance)
}
