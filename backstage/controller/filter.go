package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 过滤器控制器
 * @auth eyesYeager
 * @date 2024/1/11 17:59
 */

// UpdateAppFilter
// @Summary			更新应用过滤器
// @Description		权限：admin
// @Tags			filter
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppFilter		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/filter/updateAppFilter [post]
func UpdateAppFilter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppFilter
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.FilterService.UpdateAppFilter(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// UpdateAppLimiter
// @Summary			更新应用限流器
// @Description		权限：admin
// @Tags			filter
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppLimiter		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/filter/updateAppLimiter [post]
func UpdateAppLimiter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppLimiter
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.FilterService.UpdateAppLimiter(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// UpdateAppNameList
// @Summary			更新应用限流器
// @Description		权限：admin
// @Tags			filter
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppNameList		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/filter/updateAppNameList [post]
func UpdateAppNameList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppNameList
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.FilterService.UpdateAppNameList(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// AddNameListIp
// @Summary			添加黑白名单ip
// @Description		权限：admin
// @Tags			filter
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.AddNameListIp			body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/filter/addNameListIp [post]
func AddNameListIp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.AddNameListIp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.FilterService.AddNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// DeleteNameListIp
// @Summary			删除黑白名单ip
// @Description		权限：admin
// @Tags			filter
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.DeleteNameListIp		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/filter/deleteNameListIp [post]
func DeleteNameListIp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.DeleteNameListIp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.FilterService.DeleteNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
