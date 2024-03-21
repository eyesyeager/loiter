package service

import (
	"loiter/backstage/constant"
	"loiter/kernel/container"
	"loiter/utils"
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

// RefreshAppContainer 刷新应用容器
func (*containerService) RefreshAppContainer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新应用类型容器
	if err := container.RefreshAppGenre(appId); err != nil {
		return err
	}
	// 刷新应用实例容器
	if err := container.RefreshAppServer(appId); err != nil {
		return err
	}
	// 刷新应用静态配置容器
	if err := container.RefreshAppStatic(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "应用容器"))
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

// RefreshProcessor 刷新处理器容器
func (*containerService) RefreshProcessor(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshProcessor(appId); err != nil {
		return err
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, appId,
		constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, "处理器容器"))
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
