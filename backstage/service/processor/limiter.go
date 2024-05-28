package processor

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/service"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"time"
)

/**
 * 限制器业务层
 * @auth eyesYeager
 * @date 2024/2/29 17:54
 */

type limiterService struct {
}

var LimiterService = limiterService{}

// SaveAppLimiter 更新应用限流器
func (*limiterService) SaveAppLimiter(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.SaveAppLimiter) error {
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
			Mode:      data.Mode,
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
		Mode:      data.Mode,
		Parameter: data.Parameter,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}

	// 记录操作日志
	go service.LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppLimiter, checkAppLimiter.Limiter, checkAppLimiter.Parameter, data.Limiter, data.Parameter))
	return nil
}

// DeleteAppLimiter 删除应用限流器
func (*limiterService) DeleteAppLimiter(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.DeleteAppLimiter) error {
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, data.AppId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("不存在id为%d的应用", data.AppId))
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 删除数据库数据
	if err := global.MDB.Where(&entity.AppLimiter{AppId: data.AppId}).Unscoped().Delete(&entity.AppLimiter{}).Error; err != nil {
		return err
	}
	// 删除容器数据
	container.DeleteLimiter(checkApp.Host)
	// 记录操作日志
	go service.LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.DeleteAppLimiter))
	return nil
}

// GetLimiterByPage 分页获取限流器信息
func (*limiterService) GetLimiterByPage(data receiver.GetLimiterByPage) (err error, res returnee.GetLimiterByPage) {
	res.PageStruct = data.PageStruct
	// 获取明细信息列表
	var resInnerPOList []po.GetLimiterByPage
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = global.MDB.Raw(`SELECT a.id AppId, a.name AppName, al.mode, al.limiter LimiterCode, l.name LimiterName, al.parameter, al.updated_at
       			FROM app a, app_limiter al, limiter l
         		WHERE a.id = al.app_id AND al.limiter = l.code 
         		  AND (0 = ? OR a.id = ?) AND ('' = ? OR al.limiter = ?)
				ORDER BY a.updated_at DESC
				LIMIT ?, ?`, data.AppId, data.AppId, data.Limiter, data.Limiter, offset, limit).Scan(&resInnerPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	if resInnerPOList == nil || len(resInnerPOList) == 0 {
		return err, res
	}
	// 处理时间格式
	var resInnerList []returnee.GetLimiterByPageInner
	for _, item := range resInnerPOList {
		var resInner returnee.GetLimiterByPageInner
		_ = copier.Copy(&resInner, &item)
		resInner.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
		resInnerList = append(resInnerList, resInner)
	}
	// 查询总数
	var total int64
	if len(resInnerList) < data.PageSize {
		total = int64(len(resInnerList))
	} else {
		if err = global.MDB.Raw(`SELECT COUNT(*) FROM app a, app_limiter al 
                WHERE a.id = al.app_id AND (? = 0 OR a.id = ?) AND (? = '' OR al.limiter = ?)`,
			data.AppId, data.AppId, data.Limiter, data.Limiter).Scan(&total).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
		}
	}
	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = resInnerList
	return err, res
}
