package service

import (
	"errors"
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
		return errors.New(err.Error())
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAllContainer()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshAllContainer, app.Name))
	}()
	return nil
}

// RefreshAppServer 刷新应用与实例容器
func (*containerService) RefreshAppServer(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshAppServer(appId); err != nil {
		return errors.New(err.Error())
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAppServer()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshAppServerContainer, app.Name))
	}()
	return nil
}

// RefreshBalance 刷新负载均衡容器
func (*containerService) RefreshBalance(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 刷新容器
	if err := container.RefreshBalance(appId); err != nil {
		return errors.New(err.Error())
	}
	// 记录操作日志
	go func() {
		app := entity.App{}
		if err := global.MDB.Where(&entity.App{Model: gorm.Model{ID: appId}}).First(&app).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshBalance()-doLog", err.Error()))
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.RefreshBalanceContainer, app.Name))
	}()
	return nil
}
