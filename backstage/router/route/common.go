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
	routerRoot.GET(group+"/getNoticeDictionary", controller.GetNoticeDictionary)
	routerRoot.GET(group+"/getRoleDictionary", controller.GetRoleDictionary)
	routerRoot.GET(group+"/getProcessorDictionary", controller.GetProcessorDictionary)
}
