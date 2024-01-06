package service

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/utils"
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
func (*balanceService) AddAppBalance() {

}

// UpdateAppBalance 更新应用负载均衡策略
func (*balanceService) UpdateAppBalance(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppBalance) error {
	// 获取应用当前策略
	var checkAppBalance entity.AppBalance
	checkAppBalanceTX := global.MDB.Where(&entity.AppBalance{AppId: data.AppID}).First(&checkAppBalance)
	if checkAppBalanceTX.Error != nil {
		errMsg := fmt.Sprintf(result.ResultInfo.DbOperateError, checkAppBalanceTX.Error.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}

	// 如果新策略与当前策略相同，则取消更新
	if checkAppBalanceTX.RowsAffected != 0 && checkAppBalance.BalanceId == data.BalanceId {
		return errors.New("更新策略与当前策略相同！")
	}

	// 校验是否存在对应策略
	var balanceSlice []entity.Balance
	if err := global.MDB.Find(&balanceSlice, []uint{checkAppBalance.BalanceId, data.BalanceId}).Error; err != nil {
		errMsg := fmt.Sprintf(result.ResultInfo.DbOperateError, checkAppBalanceTX.Error.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}
	if len(balanceSlice) != 2 {
		return errors.New(fmt.Sprintf("非法负载均衡策略id！[%s, %s]中存在无效BalanceId!"))
	}

	// Slice转Map

	if checkAppBalanceTX.RowsAffected == 0 {
		// 如果对应应用没有配置策略，则添加策略
		//global.MDB.Create()
	} else {
		// 如果对应应用已经配置了策略，则修改策略
	}

	// 记录操作日志
	//go LogService.Universal(r, userClaims.Uid,
	//	constant.BuildUniversalLog(constant.LogUniversal.AddApp, data.Name, data.Host, data.Remarks))
	return nil
}
