package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/model/receiver"
	"loiter/utils"
	"net/http"
	"strconv"
)

/**
 * 应用控制器
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// SaveApp
// @Summary			注册/编辑应用
// @Description		权限：admin
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			receiver.SaveApp		body		string		false		"请求参数"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/app/saveApp [post]
func SaveApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.SaveApp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = service.AppService.SaveApp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// ActivateApp
// @Summary			激活/失效应用
// @Description		权限：admin
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			receiver.ActivateApp		body		string		false		"请求参数"
// @Param			token						header		string		true		"身份令牌"
// @Success			200							{object}	result.Response
// @Failure			400							{object}	result.Response
// @Router			/app/activateApp [post]
func ActivateApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.ActivateApp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = service.AppService.ActivateApp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// DeleteApp
// @Summary			删除应用
// @Description		权限：admin
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			receiver.DeleteApp		body		string		false		"请求参数"
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/app/deleteApp [post]
func DeleteApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.DeleteApp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = service.AppService.DeleteApp(r, userClaims, data.AppId); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAllApp
// @Summary			分页获取应用
// @Description		权限：user
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/app/getAllApp [get]
func GetAllApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.AppService.GetAllApp(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAppInfoByPage
// @Summary			获取所有应用
// @Description		权限：user
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetAppInfoByPage			body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/app/getAppInfoByPage [post]
func GetAppInfoByPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetAppInfoByPage
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 分页处理
	if err := utils.CheckPageStruct(data.PageStruct); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := service.AppService.GetAppInfoByPage(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAppInfoById
// @Summary			根据id获取应用信息
// @Description		权限：user
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/app/getAppInfoById/:appId [get]
func GetAppInfoById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验格式转换
	appId, err := strconv.Atoi(p.ByName("appId"))
	if err != nil {
		result.FailAttachedMsg(w, fmt.Sprintf("appId格式错误，error：%s", err.Error()))
		return
	}
	if appId <= 0 {
		result.FailAttachedMsg(w, fmt.Sprintf("非法参数！appId：%d", appId))
		return
	}

	// 执行业务
	if err, res := service.AppService.GetAppInfoById(uint(appId)); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
