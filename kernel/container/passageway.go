package container

import (
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/model/po"
	"strings"
)

/**
 * 通道容器
 * @auth eyesYeager
 * @date 2024/1/9 18:26
 */

// PassagewayByAppMap 通道有序列表 by AppHost
var PassagewayByAppMap map[string][]string

// InitPassageway 初始化通道容器
func InitPassageway() {
	global.AppLogger.Info("start initializing the Passageway container")
	var appPassagewayName []po.GetAppPassagewayName
	if rowsAffected := global.MDB.Raw(`SELECT a.host, ap.passageway_name 
					FROM app a, app_passageway ap 
					WHERE a.status = ? AND a.id = ap.app_id`, constant.Status.Normal).Scan(&appPassagewayName).RowsAffected; rowsAffected == 0 {
		global.AppLogger.Warn("there is currently no valid passageway configuration")
	}
	// 构建并刷新容器
	var containerMap = make(map[string][]string)
	for _, item := range appPassagewayName {
		if item.PassagewayName == "" {
			continue
		}
		nameSlice := strings.Split(item.PassagewayName, ",")
		containerMap[item.Host] = nameSlice
	}
	PassagewayByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Passageway container")
}

// RefreshPassageway 刷新通道容器
func RefreshPassageway(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Passageway container under the application with appId %d", appId))
	var appPassagewayName po.GetAppPassagewayName
	rowsAffected := global.MDB.Raw(`SELECT a.host, ap.passageway_name 
					FROM app a, app_passageway ap 
					WHERE a.status = ? AND a.app_id = ? AND a.id = ap.app_id`, constant.Status.Normal, appId).Scan(&appPassagewayName).RowsAffected
	if rowsAffected == 0 || appPassagewayName.PassagewayName == "" {
		delete(PassagewayByAppMap, appPassagewayName.Host)
	}
	PassagewayByAppMap[appPassagewayName.Host] = strings.Split(appPassagewayName.PassagewayName, ",")
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Passageway container under the application with appId %d", appId))
	return nil
}
