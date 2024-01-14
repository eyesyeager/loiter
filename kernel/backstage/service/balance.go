package service

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/utils"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 负载均衡业务层
 * @auth eyesYeager
 * @date 2024/1/5 16:51
 */

type balanceService struct {
}

var BalanceService = balanceService{}

// AddAppBalance 添加应用负载均衡策略
func (*balanceService) AddAppBalance(appId uint, balanceId uint) error {
	if err := global.MDB.Create(&entity.AppBalance{
		AppId:     appId,
		BalanceId: balanceId,
	}).Error; err != nil {
		errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "AddAppBalance()-Create", err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// UpdateAppBalance 更新应用负载均衡策略
func (*balanceService) UpdateAppBalance(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppBalance) error {
	// 获取应用当前策略
	var checkAppBalance entity.AppBalance
	checkAppBalanceTX := global.MDB.Where(&entity.AppBalance{AppId: data.AppID}).First(&checkAppBalance)

	// 如果新策略与当前策略相同，则取消更新
	if checkAppBalanceTX.RowsAffected != 0 && checkAppBalance.BalanceId == data.BalanceId {
		return errors.New("更新策略与当前策略相同！")
	}

	// 校验是否存在对应策略
	var balanceIdSlice []uint
	if checkAppBalanceTX.RowsAffected == 0 {
		balanceIdSlice = []uint{data.BalanceId}
	} else {
		balanceIdSlice = []uint{checkAppBalance.BalanceId, data.BalanceId}
	}
	var balanceSlice []entity.Balance
	if err := global.MDB.Find(&balanceSlice, balanceIdSlice).Error; err != nil {
		errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppBalance()-balanceSlice", err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}
	if checkAppBalanceTX.RowsAffected == 0 && len(balanceSlice) < 1 {
		return errors.New(fmt.Sprintf("非法负载均衡策略id！%d是无效BalanceId!", data.BalanceId))
	} else if checkAppBalanceTX.RowsAffected != 0 && len(balanceSlice) < 2 {
		return errors.New(fmt.Sprintf("非法负载均衡策略id！[%d(原id), %d(新id)]中存在无效BalanceId!", checkAppBalance.BalanceId, data.BalanceId))
	}

	// 如果原应用未配置策略，则直接插入
	if checkAppBalanceTX.RowsAffected == 0 {
		if err := global.MDB.Create(&entity.AppBalance{
			AppId:     data.AppID,
			BalanceId: data.BalanceId,
		}).Error; err != nil {
			errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppBalance()-checkAppBalanceTX", err.Error())
			global.BackstageLogger.Error(errMsg)
			return errors.New(errMsg)
		}
		// 记录操作日志
		go func() {
			var app entity.App
			if err := global.MDB.First(&app, data.AppID).Error; err != nil {
				global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppBalance()-doLog-insert", err.Error()))
			}
			LogService.Universal(r, userClaims.Uid,
				constant.BuildUniversalLog(constant.LogUniversal.UpdateAppBalance, app.Name, "", balanceSlice[0].Name))
		}()
		return nil
	}

	// 更新应用策略
	if err := global.MDB.Model(&entity.AppBalance{}).Where("app_id", data.AppID).Update("balance_id", data.BalanceId).Error; err != nil {
		errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppBalance()-Update", err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}

	// 记录操作日志
	go func() {
		var app entity.App
		if err := global.MDB.First(&app, data.AppID).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppBalance()-doLog-update", err.Error()))
		}
		balanceEntityById := make(map[uint]entity.Balance)
		for _, item := range balanceSlice {
			balanceEntityById[item.ID] = item
		}
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppBalance, app.Name, balanceEntityById[checkAppBalance.BalanceId].Name, balanceEntityById[data.BalanceId].Name))
	}()
	return nil
}
