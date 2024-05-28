package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"net/http"
)

/**
 * 通用数据控制器
 * @auth eyesYeager
 * @date 2024/2/22 11:54
 */

// GetStatusDictionary
// @Summary			获取状态字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getStatusDictionary [get]
func GetStatusDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetStatusDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAppDictionary
// @Summary			获取应用字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getAppDictionary [get]
func GetAppDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetAppDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetBalancerDictionary
// @Summary			获取负载均衡字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getBalancerDictionary [get]
func GetBalancerDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetBalancerDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetNoticeDictionary
// @Summary			获取状态字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getNoticeDictionary [get]
func GetNoticeDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetNoticeDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetRoleDictionary
// @Summary			获取角色字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getRoleDictionary [get]
func GetRoleDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetRoleDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetProcessorDictionary
// @Summary			获取处理器字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getProcessorDictionary [get]
func GetProcessorDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetProcessorDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAppGenreDictionary
// @Summary			获取应用类型字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getAppGenreDictionary [get]
func GetAppGenreDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetAppGenreDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetLimiterDictionary
// @Summary			获取应用限流器字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getLimiterDictionary [get]
func GetLimiterDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetLimiterDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetLimiterModeDictionary
// @Summary			获取限流器模式字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getLimiterModeDictionary [get]
func GetLimiterModeDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetLimiterModeDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetNameListDictionary
// @Summary			获取黑白名单字典
// @Description		权限：user
// @Tags			common
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/common/getNameListDictionary [get]
func GetNameListDictionary(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.CommonService.GetNameListDictionary(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
