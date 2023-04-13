package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"zliway/kernel/backstage/controller/helper"
	"zliway/kernel/backstage/controller/result"
	"zliway/kernel/backstage/controller/validator"
	"zliway/kernel/backstage/model/receiver"
	"zliway/kernel/backstage/service"
)

/**
 * 注册应用相关接口
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// AddApp 注册应用
func AddApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if !helper.CheckAuth(r) {
		result.FailAttachedMsg(w, r, "identity authentication failed")
		return
	}

	// 参数校验
	var data receiver.AppAdd
	_ = helper.PostData(r, &data)
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err := service.AppService.AddApp(data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
