package service

import (
	"loiter/backstage/constant"
	"loiter/backstage/utils"
	"loiter/kernel/container"
	"net/http"
)

/**
 * 注册容器业务层
 * @auth eyesYeager
 * @date 2024/1/8 19:58
 */

type containerService struct {
}

var ContainerService = containerService{}

// RefreshAllContainer 刷新所有容器
func (*containerService) RefreshAllContainer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshRegister(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "所有容器"))
	return nil
}

// RefreshAppServer 刷新应用与实例容器
func (*containerService) RefreshAppServer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshAppServer(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "应用与实例容器"))
	return nil
}

// RefreshBalancer 刷新负载均衡容器
func (*containerService) RefreshBalancer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshBalancer(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "负载均衡容器"))
	return nil
}

// RefreshFilter 刷新过滤器容器
func (*containerService) RefreshFilter(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshFilter(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "过滤器容器"))
	return nil
}

// RefreshLimiter 刷新限流器容器
func (*containerService) RefreshLimiter(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshLimiter(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "限流器容器"))
	return nil
}

// RefreshNameList 刷新黑白名单容器
func (*containerService) RefreshNameList(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshNameList(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "黑白名单容器"))
	return nil
}

// RefreshAid 刷新响应处理器容器
func (*containerService) RefreshAid(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshAid(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "响应处理器容器"))
	return nil
}
