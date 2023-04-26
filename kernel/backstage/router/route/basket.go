package route

import (
	"github.com/julienschmidt/httprouter"
	"zliway/kernel/backstage/controller"
)

/**
 * basket相关路由
 * @author eyesYeager
 * @date 2023/4/26 16:54
 */

func InitBasketRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/addBasket", controller.AddBasket)
}
