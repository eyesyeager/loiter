package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 日志相关路由
 * @auth eyesYeager
 * @date 2024/1/3 19:15
 */

func InitLogRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/getLoginLog", controller.GetLoginLog)
	routerRoot.POST(group+"/getUniversalLog", controller.GetUniversalLog)
}
