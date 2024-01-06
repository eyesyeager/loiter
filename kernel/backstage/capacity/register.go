package capacity

import (
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/container"
)

/**
 * 注册信息能力
 * @author eyesYeager
 * @date 2023/11/26 16:33
 */

// InitRegister 初始化/刷新所有注册信息
func InitRegister() {
	global.BackstageLogger.Info("start initializing or refreshing all container")
	InitAppServer()
	InitBalance()
}

// InitAppServer 初始化/刷新应用与实例容器
func InitAppServer() {
	global.BackstageLogger.Info("start initializing or refreshing the AppServer container")
	// 获取有效应用实例
	var serverSlice []entity.Server
	if tx := global.MDB.Where(&entity.Server{Status: constant.Status.Normal}).Find(&serverSlice); tx.RowsAffected == 0 {
		global.BackstageLogger.Warn("there are currently no valid server requiring service")
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
		global.BackstageLogger.Warn("there are currently no valid app requiring service")
		return
	}

	// 构建并刷新容器
	containerMap := make(map[string][]container.ServerWeight)
	for _, app := range appSlice {
		var tempServerSlice []container.ServerWeight
		for _, server := range serverSlice {
			if server.AppId == app.ID {
				tempServerSlice = append(tempServerSlice, container.ServerWeight{
					Server: server.Address,
					Weight: server.Weight,
				})
			}
		}
		containerMap[app.Host] = tempServerSlice
	}
	container.ServerByAppMap = containerMap
}

// InitBalance 初始化/刷新负载均衡容器
func InitBalance() {
	global.BackstageLogger.Info("start initializing or refreshing the Balance container")
}
