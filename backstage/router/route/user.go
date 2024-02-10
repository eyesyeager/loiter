package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 用户相关路由
 * @author eyesYeager
 * @date 2023/9/27 8:59
 */

func InitUserRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/doRegister", controller.DoRegister)
	routerRoot.POST(group+"/doLogin", controller.DoLogin)
}
