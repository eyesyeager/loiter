package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 通知相关路由
 * @auth eyesYeager
 * @date 2024/2/23 11:22
 */

func InitNoticeRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/getNoticeList", controller.GetNoticeList)
	routerRoot.GET(group+"/getEmailNoticeContent/:id", controller.GetEmailNoticeContent)
}
