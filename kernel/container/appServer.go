package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/global"
	"loiter/model/entity"
)

/**
 * 应用与应用实例容器
 * @auth eyesYeager
 * @date 2024/1/5 15:11
 */

// ServerByAppMap 应用实例 by AppHost
var ServerByAppMap = make(map[string][]ServerWeight)

type ServerWeight struct {
	Server string
	Weight uint
}

// InitAppServer 初始化应用与实例容器
func InitAppServer() {
	global.AppLogger.Info("start initializing the AppServer container")
	// 获取有效应用实例
	var serverList []entity.Server
	if tx := global.MDB.Where(&entity.Server{Status: constant.Status.Normal.Code}).Find(&serverList); tx.RowsAffected == 0 {
		global.AppLogger.Warn("there are currently no valid server requiring service")
		return
	}
	// 应用id去重
	var appIdList []uint
	tempMap := make(map[uint]struct{})
	for _, item := range serverList {
		if _, ok := tempMap[item.AppId]; !ok {
			tempMap[item.AppId] = struct{}{}
			appIdList = append(appIdList, item.AppId)
		}
	}
	// 获取有效应用
	var appList []entity.App
	if tx := global.MDB.Where(&entity.App{Status: constant.Status.Normal.Code}).Find(&appList, appIdList); tx.RowsAffected == 0 {
		global.AppLogger.Warn("there are currently no valid app requiring service")
		return
	}

	// 构建并刷新容器
	containerMap := make(map[string][]ServerWeight)
	for _, app := range appList {
		var tempServerList []ServerWeight
		for _, server := range serverList {
			if server.AppId == app.ID {
				tempServerList = append(tempServerList, ServerWeight{
					Server: server.Address,
					Weight: server.Weight,
				})
			}
		}
		containerMap[app.Host] = tempServerList
	}
	ServerByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of AppServer container")
}

// RefreshAppServer 刷新应用与实例容器
func RefreshAppServer(appId uint) error {
	// 获取对应应用
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新应用与实例容器失败，应用不存在！", appId))
	}
	// 若状态为无效，则删除
	if checkApp.Status != constant.Status.Normal.Code {
		delete(ServerByAppMap, checkApp.Host)
		return nil
	}
	// 获取有效应用实例
	var serverList []entity.Server
	if tx := global.MDB.Where(&entity.Server{Status: constant.Status.Normal.Code, AppId: appId}).Find(&serverList); tx.RowsAffected == 0 {
		global.AppLogger.Warn(fmt.Sprintf("appId为%d的应用刷新应用与实例容器失败，当前应用不存在有效实例！", appId))
		delete(ServerByAppMap, checkApp.Host)
		return nil
	}
	// 构建并刷新容器
	var currentServerList []ServerWeight
	for _, server := range serverList {
		if server.AppId == checkApp.ID {
			currentServerList = append(currentServerList, ServerWeight{
				Server: server.Address,
				Weight: server.Weight,
			})
		}
	}
	ServerByAppMap[checkApp.Host] = currentServerList
	return nil
}

// DeleteAppServer 删除应用与实例容器项
func DeleteAppServer(host string) {
	delete(ServerByAppMap, host)
}
