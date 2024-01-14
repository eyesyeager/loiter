package container

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/model/po"
)

/**
 * 负载均衡容器
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// BalanceByAppMap 负载均衡策略 by AppHost
var BalanceByAppMap map[string]string

// InitBalance 初始化负载均衡容器
func InitBalance() {
	global.AppLogger.Info("start initializing the Balance container")
	// 获取有效应用负载策略
	var appBalanceNameSlice []po.GetAppBalanceName
	if affected := global.MDB.Raw(`SELECT a.host, b.name 
						FROM app a, balance b, app_balance ab 
						WHERE a.status = ? AND a.id = ab.app_id AND ab.balance_id = b.id`,
		constant.Status.Normal).Scan(&appBalanceNameSlice).RowsAffected; affected == 0 {
		global.AppLogger.Warn("there is currently no valid load balancing configuration")
	}

	// 构建并刷新容器
	var containerMap = make(map[string]string)
	for _, item := range appBalanceNameSlice {
		containerMap[item.Host] = item.Name
	}
	BalanceByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Balance container")
}

// RefreshBalance 刷新负载均衡容器
func RefreshBalance(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Balance container under the application with appId %d", appId))
	// 获取有效应用负载策略
	var appBalanceName po.GetAppBalanceName
	if affected := global.MDB.Raw(`SELECT a.host, b.name 
						FROM app a, balance b, app_balance ab 
						WHERE a.id = ? AND a.status = ? AND a.id = ab.app_id AND ab.balance_id = b.id`,
		appId, constant.Status.Normal).Scan(&appBalanceName).RowsAffected; affected == 0 {
		errMsg := fmt.Sprintf("there is currently no valid load balancing configuration under the application with appId %d", appId)
		global.AppLogger.Warn(errMsg)
		return errors.New(errMsg)
	}
	BalanceByAppMap[appBalanceName.Host] = appBalanceName.Name
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Balance container under the application with appId %d", appId))
	return nil
}
