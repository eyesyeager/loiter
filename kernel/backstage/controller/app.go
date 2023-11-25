package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/9/26 14:42
 */
/**
 * app控制器
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// AddApp 注册应用
func AddApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//// 权限校验
	//if !helper.CheckAuth(r) {
	//	result.FailByCustom(w, r, result.Results.AuthError)
	//	return
	//}
	//
	//// 参数校验
	//var data receiver.AppAdd
	//if err := helper.PostData(r, &data); err != nil {
	//	result.FailAttachedMsg(w, r, err.Error())
	//	return
	//}
	//if err := validator.Checker.Struct(data); err != nil {
	//	result.FailAttachedMsg(w, r, err.Error())
	//	return
	//}
	//
	//// 执行业务
	//if err := service.AppService.AddApp(r, data); err == nil {
	//	result.SuccessDefault(w, r, nil)
	//} else {
	//	result.FailAttachedMsg(w, r, err.Error())
	//}
}
