package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/service"
	"loiter/kernel/backstage/utils"
	receiver2 "loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 日志控制器
 * @auth eyesYeager
 * @date 2024/1/3 19:13
 */

// GetLoginLog
// @Summary			登录日志查询
// @Description		权限：user
// @Tags			log
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			receiver.GetLoginLog	body		string		false		"请求参数"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/log/getLoginLog [post]
func GetLoginLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver2.GetLoginLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 分页处理
	if err := utils.CheckPageStruct(data.PageStruct); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err, res := service.LogService.GetLoginLog(data); err == nil {
		result.SuccessDefault(w, r, res)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// GetUniversalLog
// @Summary			通用日志查询
// @Description		权限：user
// @Tags			log
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token						header		string		true		"身份令牌"
// @Param			receiver.GetUniversalLog	body		string		false		"请求参数"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/log/getUniversalLog [post]
func GetUniversalLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver2.GetUniversalLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 分页处理
	if err := utils.CheckPageStruct(data.PageStruct); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err, res := service.LogService.GetUniversalLog(data); err == nil {
		result.SuccessDefault(w, r, res)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
