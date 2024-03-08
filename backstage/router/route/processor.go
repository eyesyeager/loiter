package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * 处理器路由
 * @auth eyesYeager
 * @date 2024/1/11 18:53
 */

func InitProcessorRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.POST(group+"/saveAppProcessor", controller.SaveAppProcessor)
	routerRoot.POST(group+"/getProcessorByPage", controller.GetProcessorByPage)
	routerRoot.GET(group+"/getProcessorByGenre/:genre", controller.GetProcessorByGenre)
	routerRoot.POST(group+"/updateAppLimiter", controller.UpdateAppLimiter)
	routerRoot.POST(group+"/updateAppNameList", controller.UpdateAppNameList)
	routerRoot.POST(group+"/addNameListIp", controller.AddNameListIp)
	routerRoot.POST(group+"/deleteNameListIp", controller.DeleteNameListIp)
	routerRoot.GET(group+"/getOverviewRequestLog", controller.GetOverviewRequestLog)
	routerRoot.POST(group+"/getDetailedRequestExtremumLog", controller.GetDetailedRequestExtremumLog)
	routerRoot.POST(group+"/getDetailedRequestNumLog", controller.GetDetailedRequestNumLog)
	routerRoot.POST(group+"/getDetailedRequestRuntimeLog", controller.GetDetailedRequestRuntimeLog)
	routerRoot.POST(group+"/getDetailedRequestQPSLog", controller.GetDetailedRequestQPSLog)
	routerRoot.POST(group+"/getDetailedRequestVisitorLog", controller.GetDetailedRequestVisitorLog)
	routerRoot.POST(group+"/getDetailedRequestTopApiLog", controller.GetDetailedRequestTopApiLog)
	routerRoot.POST(group+"/getDetailedRequestRejectLog", controller.GetDetailedRequestRejectLog)
}
