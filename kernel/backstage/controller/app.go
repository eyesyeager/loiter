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
 * 应用相关控制器
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// AddApp 注册应用
func AddApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if !helper.CheckAuth(r) {
		result.FailByCustom(w, r, result.Results.AuthError)
		return
	}

	// 参数校验
	var data receiver.AppAdd
	if err := helper.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err := service.AppService.AddApp(r, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// AddAppServer 给应用添加服务
func AddAppServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if !helper.CheckAuth(r) {
		result.FailByCustom(w, r, result.Results.AuthError)
		return
	}

	// 参数校验
	var data receiver.AppServerAdd
	if err := helper.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err := service.AppService.AddServerApp(r, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// GetAppAndServer 获取应用以及下属服务
func GetAppAndServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if !helper.CheckAuth(r) {
		result.FailByCustom(w, r, result.Results.AuthError)
		return
	}

	// 执行业务
	if err, data := service.AppService.GetServerAndApp(); err == nil {
		result.SuccessDefault(w, r, data)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
