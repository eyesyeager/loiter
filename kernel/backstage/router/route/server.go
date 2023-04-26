package route

import (
	"github.com/julienschmidt/httprouter"
	"zliway/kernel/backstage/controller"
)

/**
 * server相关路由
 * @author eyesYeager
 * @date 2023/4/26 16:55
 */

func InitServerRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/addServer", controller.AddServer)
}
