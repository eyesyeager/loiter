package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/controller"
)

/**
 * 应用实例相关路由
 * @auth eyesYeager
 * @date 2024/1/5 14:06
 */

func InitServerRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/addServer", controller.AddServer)
}
