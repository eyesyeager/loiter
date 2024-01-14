package container

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/model/entity"
)

/**
 * 应用与应用实例容器
 * @auth eyesYeager
 * @date 2024/1/5 15:11
 */

// ServerByAppMap 应用实例 by AppHost
var ServerByAppMap map[string][]ServerWeight

type ServerWeight struct {
	Server string
	Weight uint
}

// InitAppServer 初始化应用与实例容器
func InitAppServer() {
	global.AppLogger.Info("start initializing the AppServer container")
	// 获取有效应用实例
	var serverSlice []entity.Server
	if tx := global.MDB.Where(&entity.Server{Status: constant.Status.Normal}).Find(&serverSlice); tx.RowsAffected == 0 {
		global.AppLogger.Warn("there are currently no valid server requiring service")
		return
	}
	// 应用id去重
	var appIdSlice []uint
	tempMap := make(map[uint]struct{})
	for _, item := range serverSlice {
		if _, ok := tempMap[item.AppId]; !ok {
			tempMap[item.AppId] = struct{}{}
			appIdSlice = append(appIdSlice, item.AppId)
		}
	}
	// 获取有效应用
	var appSlice []entity.App
	if tx := global.MDB.Where(&entity.App{Status: constant.Status.Normal}).Find(&appSlice, appIdSlice); tx.RowsAffected == 0 {
		global.AppLogger.Warn("there are currently no valid app requiring service")
		return
	}

	// 构建并刷新容器
	containerMap := make(map[string][]ServerWeight)
	for _, app := range appSlice {
		var tempServerSlice []ServerWeight
		for _, server := range serverSlice {
			if server.AppId == app.ID {
				tempServerSlice = append(tempServerSlice, ServerWeight{
					Server: server.Address,
					Weight: server.Weight,
				})
			}
		}
		containerMap[app.Host] = tempServerSlice
	}
	ServerByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of AppServer container")
}

// RefreshAppServer 刷新应用与实例容器
func RefreshAppServer(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the AppServer container under the application with appId %d", appId))
	// 获取对应应用
	var app entity.App
	if tx := global.MDB.Where(&entity.App{Status: constant.Status.Normal}).First(&app, appId); tx.RowsAffected == 0 {
		errMsg := fmt.Sprintf("there are currently no valid app requiring service under the application with appId %d", appId)
		global.AppLogger.Warn(errMsg)
		return errors.New(errMsg)
	}
	// 获取有效应用实例
	var serverSlice []entity.Server
	if tx := global.MDB.Where(&entity.Server{Status: constant.Status.Normal, AppId: appId}).Find(&serverSlice); tx.RowsAffected == 0 {
		global.AppLogger.Warn(fmt.Sprintf("there are currently no valid server requiring service under the application with appId %d", appId))
		delete(ServerByAppMap, app.Host)
		return nil
	}
	// 构建并刷新容器
	var currentServerSlice []ServerWeight
	for _, server := range serverSlice {
		if server.AppId == app.ID {
			currentServerSlice = append(currentServerSlice, ServerWeight{
				Server: server.Address,
				Weight: server.Weight,
			})
		}
	}
	ServerByAppMap[app.Host] = currentServerSlice
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of AppServer container under the application with appId %d", appId))
	return nil
}
