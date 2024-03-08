package processor

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/service"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/receiver"
	"loiter/utils"
	"net/http"
)

/**
 * 限制器业务层
 * @auth eyesYeager
 * @date 2024/2/29 17:54
 */

type limiterService struct {
}

var LimiterService = limiterService{}

// UpdateAppLimiter 更新应用限流器
func (*limiterService) UpdateAppLimiter(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppLimiter) error {
	// 校验传入限流器的合法性
	var count int64
	var checkLimiter = entity.Limiter{Code: data.Limiter}
	if err := global.MDB.Model(&checkLimiter).Where(&checkLimiter).Count(&count).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if count == 0 {
		return errors.New(fmt.Sprintf("不存在编码为%s的限流器", data.Limiter))
	}
	// 获取应用当前所用限流器
	var checkAppLimiter entity.AppLimiter
	checkAppLimiterTX := global.MDB.Where(&entity.AppLimiter{AppId: data.AppId}).First(&checkAppLimiter)
	// 如果原先没有配置限流器，则插入
	if checkAppLimiterTX.RowsAffected == 0 {
		if err := global.MDB.Create(&entity.AppLimiter{
			AppId:     data.AppId,
			Limiter:   data.Limiter,
			Parameter: data.Parameter,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
		// 记录操作日志
		go service.LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppLimiter, "", "", data.Limiter, data.Parameter))
		return nil
	}

	// 如果原先已经配置了限流器，则修改
	if err := global.MDB.Model(&entity.AppLimiter{}).Where("app_id", data.AppId).Updates(entity.AppLimiter{
		Limiter:   data.Limiter,
		Parameter: data.Parameter,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}

	// 记录操作日志
	go service.LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppLimiter, checkAppLimiter.Limiter, checkAppLimiter.Parameter, data.Limiter, data.Parameter))
	return nil
}
