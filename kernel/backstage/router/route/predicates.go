package route

import (
	"github.com/julienschmidt/httprouter"
	"zliway/kernel/backstage/controller"
)

/**
 * 断言路由
 * @author eyesYeager
 * @date 2023/4/26 22:33
 */

func InitPredicatesRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/addPredicates", controller.AddPredicates)
}
