package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 应用相关路由
 * @author eyesYeager
 * @date 2023/4/11 20:09
 */

func InitAppRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/addApp", controller.AddApp)
}
