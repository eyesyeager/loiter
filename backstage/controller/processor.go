package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/backstage/service/processor"
	"loiter/model/receiver"
	"loiter/utils"
	"net/http"
)

/**
 * 处理器控制器
 * @auth eyesYeager
 * @date 2024/1/11 17:59
 */

// SaveAppProcessor
// @Summary			更新应用处理器
// @Description		权限：admin
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.SaveAppProcessor		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/saveAppProcessor [post]
func SaveAppProcessor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.SaveAppProcessor
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = service.ProcessorService.SaveAppProcessor(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetProcessorByPage
// @Summary			分页获取应用处理器
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.GetProcessorByPage		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/getProcessorByPage [post]
func GetProcessorByPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetProcessorByPage
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
	if err, res := service.ProcessorService.GetProcessorByPage(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetProcessorByGenre
// @Summary			分页获取应用处理器
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/getProcessorByGenre/:genre [get]
func GetProcessorByGenre(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.ProcessorService.GetProcessorByGenre(p.ByName("genre")); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// UpdateAppLimiter
// @Summary			更新应用限流器
// @Description		权限：admin
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppLimiter		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/updateAppLimiter [post]
func UpdateAppLimiter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.UpdateAppLimiter
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = processor.LimiterService.UpdateAppLimiter(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// UpdateAppNameList
// @Summary			更新应用黑白名单
// @Description		权限：admin
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.UpdateAppNameList		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/updateAppNameList [post]
func UpdateAppNameList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.UpdateAppNameList
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = processor.NameListService.UpdateAppNameList(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// AddNameListIp
// @Summary			添加黑白名单ip
// @Description		权限：admin
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.AddNameListIp			body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/addNameListIp [post]
func AddNameListIp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.AddNameListIp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = processor.NameListService.AddNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// DeleteNameListIp
// @Summary			删除黑白名单ip
// @Description		权限：admin
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token							header		string		true		"身份令牌"
// @Param			receiver.DeleteNameListIp		body		string		false		"请求参数"
// @Success			200								{object}	result.Response
// @Failure			400								{object}	result.Response
// @Router			/processor/deleteNameListIp [post]
func DeleteNameListIp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.DeleteNameListIp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = processor.NameListService.DeleteNameListIp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetOverviewRequestLog
// @Summary			获取请求日志概览
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getOverviewRequestLog [get]
func GetOverviewRequestLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetOverviewRequestLog(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestExtremumLog
// @Summary			获取请求日志详情-极值
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestExtremumLog [post]
func GetDetailedRequestExtremumLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestExtremumLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestNumLog
// @Summary			获取请求日志详情-请求数
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestNumLog [post]
func GetDetailedRequestNumLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestNumLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestRuntimeLog
// @Summary			获取请求日志详情-响应时间
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestRuntimeLog [post]
func GetDetailedRequestRuntimeLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestRuntimeLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestQPSLog
// @Summary			获取请求日志详情-QPS
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestQPSLog [post]
func GetDetailedRequestQPSLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestQPSLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestVisitorLog
// @Summary			获取请求日志详情-访客
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestVisitorLog [post]
func GetDetailedRequestVisitorLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestVisitorLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestTopApiLog
// @Summary			获取请求日志详情-Top接口
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestTopApiLog [post]
func GetDetailedRequestTopApiLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestTopApiLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetDetailedRequestRejectLog
// @Summary			获取请求日志详情-Top接口
// @Description		权限：user
// @Tags			processor
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.GetDetailedRequestLog		body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/processor/getDetailedRequestRejectLog [post]
func GetDetailedRequestRejectLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 参数校验
	var data receiver.GetDetailedRequestLog
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err, res := processor.RequestLogService.GetDetailedRequestRejectLog(data); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
