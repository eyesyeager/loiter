package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/service"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 通道控制器
 * @auth eyesYeager
 * @date 2024/1/11 17:59
 */

// UpdateAppPassageway
// @Summary			更新应用通道
// @Description		权限：admin
// @Tags			passageway
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppPassageway	body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/passageway/updateAppPassageway [post]
func UpdateAppPassageway(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppPassageway
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.PassagewayService.UpdateAppPassageway(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// UpdateAppLimiter
// @Summary			更新应用限流器
// @Description		权限：admin
// @Tags			passageway
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppLimiter		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/passageway/updateAppLimiter [post]
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
	if err = service.PassagewayService.UpdateAppLimiter(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// UpdateAppNameList
// @Summary			更新应用限流器
// @Description		权限：admin
// @Tags			passageway
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppNameList		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/passageway/updateAppNameList [post]
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
	if err = service.PassagewayService.UpdateAppNameList(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// AddNameListIp
// @Summary			添加黑白名单ip
// @Description		权限：admin
// @Tags			passageway
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.AddNameListIp			body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/passageway/addNameListIp [post]
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
	if err = service.PassagewayService.AddNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// DeleteNameListIp
// @Summary			删除黑白名单ip
// @Description		权限：admin
// @Tags			passageway
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.DeleteNameListIp		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/passageway/deleteNameListIp [post]
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
	if err = service.PassagewayService.DeleteNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
