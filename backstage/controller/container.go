package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"net/http"
	"strconv"
)

/**
 * 注册容器控制器
 * @auth eyesYeager
 * @date 2024/1/8 19:54
 */

// RefreshAllContainer
// @Summary			刷新所有容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshAllContainer/:appId [get]
func RefreshAllContainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshAllContainer(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// RefreshAppServer
// @Summary			刷新应用与实例容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshAppServer/:appId [get]
func RefreshAppServer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshAppServer(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// RefreshBalancer
// @Summary			刷新负载均衡容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshBalancer/:appId [get]
func RefreshBalancer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshBalancer(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// RefreshProcessor
// @Summary			刷新处理器容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshProcessor/:appId [get]
func RefreshProcessor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshProcessor(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// RefreshLimiter
// @Summary			刷新限流器容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshLimiter/:appId [get]
func RefreshLimiter(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshLimiter(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// RefreshNameList
// @Summary			刷新黑白名单容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			appId					path		string		true
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshNameList/:appId [get]
func RefreshNameList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
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
	if err = service.ContainerService.RefreshNameList(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
