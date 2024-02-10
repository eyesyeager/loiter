package container

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/po"
	"strings"
)

/**
 * 过滤器容器
 * @auth eyesYeager
 * @date 2024/1/9 18:26
 */

// FilterByAppMap 过滤器有序列表 by AppHost
var FilterByAppMap map[string][]string

// InitFilter 初始化过滤器容器
func InitFilter() {
	global.AppLogger.Info("start initializing the Filter container")
	var appFilterName []po.GetAppFilterName
	if rowsAffected := global.MDB.Raw(`SELECT a.host, ap.filter_name 
						FROM app a, app_filter ap 
						WHERE a.status = ? AND a.id = ap.app_id`, constant.Status.Normal).Scan(&appFilterName).RowsAffected; rowsAffected == 0 {
		global.AppLogger.Info("there is currently no valid Filter configuration")
		return
	}
	// 构建并刷新容器
	var containerMap = make(map[string][]string)
	for _, item := range appFilterName {
		if item.FilterName == "" {
			continue
		}
		nameSlice := strings.Split(item.FilterName, config.Program.PluginConfig.FilterDelimiter)
		containerMap[item.Host] = nameSlice
	}
	FilterByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Filter container")
}

// RefreshFilter 刷新过滤器容器
func RefreshFilter(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Filter container under the application with appId %d", appId))
	var appFilterName po.GetAppFilterName
	tx := global.MDB.Raw(`SELECT a.host, ap.filter_name 
					FROM app a, app_filter ap 
					WHERE a.status = ? AND a.id = ? AND a.id = ap.app_id`, constant.Status.Normal, appId).Scan(&appFilterName)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
	}
	// 查询为空则删除元素
	if tx.RowsAffected == 0 {
		var checkApp = entity.App{Model: gorm.Model{ID: appId}}
		if err := global.MDB.First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("appId为%d的应用不存在或者无效！", appId))
		}
		delete(FilterByAppMap, checkApp.Host)
		return nil
	}
	// 刷新容器
	FilterByAppMap[appFilterName.Host] = strings.Split(appFilterName.FilterName, config.Program.PluginConfig.FilterDelimiter)
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Filter container under the application with appId %d", appId))
	return nil
}
