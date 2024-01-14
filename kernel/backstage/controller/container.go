package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/service"
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
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshAllContainer/:appId [get]
func RefreshAllContainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验格式转换
	appId, err := strconv.Atoi(p.ByName("appId"))
	if err != nil {
		result.FailAttachedMsg(w, r, fmt.Sprintf("appId格式错误，error：%s", err.Error()))
		return
	}
	if appId <= 0 {
		result.FailAttachedMsg(w, r, fmt.Sprintf("非法参数！appId：%d", appId))
		return
	}

	// 执行业务
	if err = service.ContainerService.RefreshAllContainer(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
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
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshAppServer/:appId [get]
func RefreshAppServer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验格式转换
	appId, err := strconv.Atoi(p.ByName("appId"))
	if err != nil {
		result.FailAttachedMsg(w, r, fmt.Sprintf("appId格式错误，error：%s", err.Error()))
		return
	}
	if appId <= 0 {
		result.FailAttachedMsg(w, r, fmt.Sprintf("非法参数！appId：%d", appId))
		return
	}

	// 执行业务
	if err = service.ContainerService.RefreshAppServer(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// RefreshBalance
// @Summary			刷新负载均衡容器
// @Description		权限：admin
// @Tags			container
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/container/refreshBalance/:appId [get]
func RefreshBalance(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验格式转换
	appId, err := strconv.Atoi(p.ByName("appId"))
	if err != nil {
		result.FailAttachedMsg(w, r, fmt.Sprintf("appId格式错误，error：%s", err.Error()))
		return
	}
	if appId <= 0 {
		result.FailAttachedMsg(w, r, fmt.Sprintf("非法参数！appId：%d", appId))
		return
	}

	// 执行业务
	if err = service.ContainerService.RefreshBalance(r, userClaims, uint(appId)); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
