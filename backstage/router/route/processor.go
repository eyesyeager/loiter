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
	// 处理器
	routerRoot.POST(group+"/saveAppProcessor", controller.SaveAppProcessor)
	routerRoot.POST(group+"/getProcessorByPage", controller.GetProcessorByPage)
	routerRoot.GET(group+"/getProcessorByGenre/:genre", controller.GetProcessorByGenre)
	// 限流器
	routerRoot.POST(group+"/saveAppLimiter", controller.SaveAppLimiter)
	routerRoot.POST(group+"/deleteAppLimiter", controller.DeleteAppLimiter)
	routerRoot.POST(group+"/getLimiterByPage", controller.GetLimiterByPage)
	// 黑白名单
	routerRoot.POST(group+"/updateAppNameList", controller.UpdateAppNameList)
	routerRoot.POST(group+"/addNameListIp", controller.AddNameListIp)
	routerRoot.POST(group+"/deleteNameListIp", controller.DeleteNameListIp)
	// 请求日志
	routerRoot.GET(group+"/getOverviewRequestLog", controller.GetOverviewRequestLog)
	routerRoot.POST(group+"/getDetailedRequestExtremumLog", controller.GetDetailedRequestExtremumLog)
	routerRoot.POST(group+"/getDetailedRequestNumLog", controller.GetDetailedRequestNumLog)
	routerRoot.POST(group+"/getDetailedRequestRuntimeLog", controller.GetDetailedRequestRuntimeLog)
	routerRoot.POST(group+"/getDetailedRequestQPSLog", controller.GetDetailedRequestQPSLog)
	routerRoot.POST(group+"/getDetailedRequestVisitorLog", controller.GetDetailedRequestVisitorLog)
	routerRoot.POST(group+"/getDetailedRequestTopApiLog", controller.GetDetailedRequestTopApiLog)
	routerRoot.POST(group+"/getDetailedRequestRejectLog", controller.GetDetailedRequestRejectLog)
}
