package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/controller"
)

/**
 * 响应处理器路由
 * @auth eyesYeager
 * @date 2024/1/26 16:11
 */

func InitAidRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/updateAppAid", controller.UpdateAppAid)
}
