package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/global"
	"loiter/kernel/model/po"
)

/**
 * 负载均衡容器
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// BalancerByAppMap 负载均衡策略 by AppHost
var BalancerByAppMap map[string]string

// InitBalancer 初始化负载均衡容器
func InitBalancer() {
	global.AppLogger.Info("start initializing the Balancer container")
	// 获取有效应用负载策略
	var appBalancerNameSlice []po.GetAppBalancerName
	if affected := global.MDB.Raw(`SELECT a.host, b.name 
					FROM app a, balancer b, app_balancer ab 
					WHERE a.status = ? AND a.id = ab.app_id AND ab.balancer_id = b.id`,
		constant.Status.Normal).Scan(&appBalancerNameSlice).RowsAffected; affected == 0 {
		global.AppLogger.Warn("there is currently no valid Balancer configuration")
		return
	}

	// 构建并刷新容器
	var containerMap = make(map[string]string)
	for _, item := range appBalancerNameSlice {
		containerMap[item.Host] = item.Name
	}
	BalancerByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Balancer container")
}

// RefreshBalancer 刷新负载均衡容器
func RefreshBalancer(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Balancer container under the application with appId %d", appId))
	// 获取有效应用负载策略
	var appBalancerName po.GetAppBalancerName
	tx := global.MDB.Raw(`SELECT a.host, b.name 
						FROM app a, balancer b, app_balancer ab 
						WHERE a.id = ? AND a.status = ? AND a.id = ab.app_id AND ab.balancer_id = b.id`, appId, constant.Status.Normal).Scan(&appBalancerName)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
	}
	// 查询为空则返回错误信息
	if tx.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新负载均衡容器失败，负载均衡配置不能为空！", appId))
	}
	// 刷新负载均衡容器
	BalancerByAppMap[appBalancerName.Host] = appBalancerName.Name
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Balancer container under the application with appId %d", appId))
	return nil
}
