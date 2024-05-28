package route

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/controller"
)

/**
 * @author eyesYeager
 * @date 2023/11/25 22:49
 */

func InitCommonRoute(routerRoot *httprouter.Router, group string) {
	routerRoot.GET(group+"/getStatusDictionary", controller.GetStatusDictionary)
	routerRoot.GET(group+"/getAppDictionary", controller.GetAppDictionary)
	routerRoot.GET(group+"/getBalancerDictionary", controller.GetBalancerDictionary)
	routerRoot.GET(group+"/getNoticeDictionary", controller.GetNoticeDictionary)
	routerRoot.GET(group+"/getRoleDictionary", controller.GetRoleDictionary)
	routerRoot.GET(group+"/getProcessorDictionary", controller.GetProcessorDictionary)
	routerRoot.GET(group+"/getAppGenreDictionary", controller.GetAppGenreDictionary)
	routerRoot.GET(group+"/getLimiterDictionary", controller.GetLimiterDictionary)
	routerRoot.GET(group+"/getLimiterModeDictionary", controller.GetLimiterModeDictionary)
	routerRoot.GET(group+"/getNameListDictionary", controller.GetNameListDictionary)
}
