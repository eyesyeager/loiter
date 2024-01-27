package service

import (
	"fmt"
	"gorm.io/gorm"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/utils"
	"loiter/kernel/container"
	"loiter/kernel/model/entity"
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
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAllContainer()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "所有容器"))
	}()
	return nil
}

// RefreshAppServer 刷新应用与实例容器
func (*containerService) RefreshAppServer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshAppServer(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAppServer()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "应用与实例容器"))
	}()
	return nil
}

// RefreshBalancer 刷新负载均衡容器
func (*containerService) RefreshBalancer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshBalancer(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshBalancer()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "负载均衡容器"))
	}()
	return nil
}

// RefreshPassageway 刷新通道容器
func (*containerService) RefreshPassageway(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshPassageway(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshPassageway()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "通道容器"))
	}()
	return nil
}

// RefreshLimiter 刷新限流器容器
func (*containerService) RefreshLimiter(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshLimiter(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshLimiter()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "限流器容器"))
	}()
	return nil
}

// RefreshNameList 刷新黑白名单容器
func (*containerService) RefreshNameList(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshNameList(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshNameList()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "黑白名单容器"))
	}()
	return nil
}

// RefreshAid 刷新响应处理器容器
func (*containerService) RefreshAid(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshAid(appId); err != nil {
		return err
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAid()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshContainer, app.Name, "响应处理器容器"))
	}()
	return nil
}
