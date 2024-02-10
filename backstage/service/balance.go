package service

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/utils"
	"loiter/global"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 负载均衡业务层
 * @auth eyesYeager
 * @date 2024/1/5 16:51
 */

type balancerService struct {
}

var BalancerService = balancerService{}

// AddAppBalancer 添加应用负载均衡策略
func (*balancerService) AddAppBalancer(appId uint, balancerId uint) error {
	if err := global.MDB.Create(&entity.AppBalancer{
		AppId:      appId,
		BalancerId: balancerId,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	return nil
}

// UpdateAppBalancer 更新应用负载均衡策略
func (*balancerService) UpdateAppBalancer(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppBalancer) error {
	// 获取应用当前策略
	var checkAppBalancer entity.AppBalancer
	checkAppBalancerTX := global.MDB.Where(&entity.AppBalancer{AppId: data.AppId}).First(&checkAppBalancer)

	// 如果新策略与当前策略相同，则取消更新
	if checkAppBalancerTX.RowsAffected != 0 && checkAppBalancer.BalancerId == data.BalancerId {
		return errors.New("更新策略与当前策略相同！")
	}

	// 校验是否存在对应策略
	var balancerIdSlice []uint
	if checkAppBalancerTX.RowsAffected == 0 {
		balancerIdSlice = []uint{data.BalancerId}
	} else {
		balancerIdSlice = []uint{checkAppBalancer.BalancerId, data.BalancerId}
	}
	var balancerSlice []entity.Balancer
	if err := global.MDB.Find(&balancerSlice, balancerIdSlice).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if checkAppBalancerTX.RowsAffected == 0 && len(balancerSlice) < 1 {
		return errors.New(fmt.Sprintf("非法负载均衡策略id！%d是无效BalancerId!", data.BalancerId))
	} else if checkAppBalancerTX.RowsAffected != 0 && len(balancerSlice) < 2 {
		return errors.New(fmt.Sprintf("非法负载均衡策略id！[%d(原id), %d(新id)]中存在无效BalancerId!", checkAppBalancer.BalancerId, data.BalancerId))
	}

	// 如果原应用未配置策略，则直接插入
	if checkAppBalancerTX.RowsAffected == 0 {
		if err := global.MDB.Create(&entity.AppBalancer{
			AppId:      data.AppId,
			BalancerId: data.BalancerId,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
		// 记录操作日志
		go LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppBalancer, "", balancerSlice[0].Name))
		return nil
	}

	// 更新应用策略
	if err := global.MDB.Model(&entity.AppBalancer{}).Where("app_id", data.AppId).Update("balancer_id", data.BalancerId).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}

	// 记录操作日志
	go func() {
		balancerEntityById := make(map[uint]entity.Balancer)
		for _, item := range balancerSlice {
			balancerEntityById[item.ID] = item
		}
		LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppBalancer, balancerEntityById[checkAppBalancer.BalancerId].Name, balancerEntityById[data.BalancerId].Name))
	}()
	return nil
}
