package container

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/model/entity"
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
		global.AppLogger.Info("there is currently no valid Passageway configuration")
		return
	}
	// 构建并刷新容器
	var containerMap = make(map[string][]string)
	for _, item := range appPassagewayName {
		if item.PassagewayName == "" {
			continue
		}
		nameSlice := strings.Split(item.PassagewayName, config.Program.PassagewayDelimiter)
		containerMap[item.Host] = nameSlice
	}
	PassagewayByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Passageway container")
}

// RefreshPassageway 刷新通道容器
func RefreshPassageway(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Passageway container under the application with appId %d", appId))
	var appPassagewayName po.GetAppPassagewayName
	tx := global.MDB.Raw(`SELECT a.host, ap.passageway_name 
					FROM app a, app_passageway ap 
					WHERE a.status = ? AND a.id = ? AND a.id = ap.app_id`, constant.Status.Normal, appId).Scan(&appPassagewayName)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshPassageway()-Raw", tx.Error.Error()))
	}
	// 查询为空则删除元素
	if tx.RowsAffected == 0 {
		var checkApp = entity.App{Model: gorm.Model{ID: appId}}
		if err := global.MDB.First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("appId为%d的应用不存在或者无效！", appId))
		}
		delete(PassagewayByAppMap, checkApp.Host)
		return nil
	}
	// 刷新容器
	PassagewayByAppMap[appPassagewayName.Host] = strings.Split(appPassagewayName.PassagewayName, config.Program.PassagewayDelimiter)
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Passageway container under the application with appId %d", appId))
	return nil
}
