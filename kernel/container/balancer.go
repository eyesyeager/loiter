package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/global"
	"loiter/model/po"
)

/**
 * 负载均衡容器
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// BalancerByAppMap 负载均衡策略 by AppHost
var BalancerByAppMap = make(map[string]string)

// InitBalancer 初始化负载均衡容器
func InitBalancer() {
	global.AppLogger.Info("start initializing the Balancer container")
	// 获取有效应用负载策略
	var appBalancerNameList []po.GetAppBalancerName
	if affected := global.MDB.Raw(`SELECT a.host, ab.balancer 
					FROM app a, app_balancer ab 
					WHERE a.status = ? AND a.id = ab.app_id`,
		constant.Status.Normal.Code).Scan(&appBalancerNameList).RowsAffected; affected == 0 {
		global.AppLogger.Warn("there is currently no valid Balancer configuration")
		return
	}

	// 构建并刷新容器
	var containerMap = make(map[string]string)
	for _, item := range appBalancerNameList {
		containerMap[item.Host] = item.Balancer
	}
	BalancerByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Balancer container")
}

// RefreshBalancer 刷新负载均衡容器
func RefreshBalancer(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Balancer container under the application with appId %d", appId))
	// 获取有效应用负载策略
	var appBalancerName po.GetAppBalancerName
	if tx := global.MDB.Raw(`SELECT a.host, ab.balancer, a.status
						FROM app a, app_balancer ab 
						WHERE a.id = ? AND a.id = ab.app_id`, appId).Scan(&appBalancerName); tx.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新负载均衡容器失败，应用不存在！", appId))
	}
	// 状态不为normal，则删除
	if appBalancerName.Status != constant.Status.Normal.Code {
		delete(BalancerByAppMap, appBalancerName.Host)
		return nil
	}
	// 刷新负载均衡容器
	BalancerByAppMap[appBalancerName.Host] = appBalancerName.Balancer
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Balancer container under the application with appId %d", appId))
	return nil
}

// DeleteBalancer 删除负载均衡容器项
func DeleteBalancer(host string) {
	delete(BalancerByAppMap, host)
}
