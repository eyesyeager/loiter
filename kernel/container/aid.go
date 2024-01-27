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
 * 响应处理器
 * @auth eyesYeager
 * @date 2024/1/26 12:04
 */

// AidByAppMap 响应处理器 by AppHost
var AidByAppMap map[string][]string

// InitAid 初始化响应处理器容器
func InitAid() {
	global.AppLogger.Info("start initializing the Aid container")
	// 获取响应处理器配置
	var appAidName []po.GetAppAidName
	if affected := global.MDB.Raw(`SELECT a.host, aa.aid_name 
					FROM app a, app_aid aa 
					WHERE a.status = ? AND a.id = aa.app_id`, constant.Status.Normal).Scan(&appAidName).RowsAffected; affected == 0 {
		global.AppLogger.Info("there is currently no valid Aid configuration")
		return
	}
	// 构建响应处理器容器
	var containerMap = make(map[string][]string)
	for _, item := range appAidName {
		if item.AidName == "" {
			continue
		}
		nameSlice := strings.Split(item.AidName, config.Program.AidDelimiter)
		containerMap[item.Host] = nameSlice
	}
	AidByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Aid container")
}

// RefreshAid 刷新响应处理器容器
func RefreshAid(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start initializing the Aid container under the application with appId %d", appId))
	var appAidName po.GetAppAidName
	// 获取响应处理器配置
	tx := global.MDB.Raw(`SELECT a.host, aa.aid_name 
					FROM app a, app_aid aa 
					WHERE a.id = ? AND a.status = ? AND a.id = aa.app_id`, appId, constant.Status.Normal).Scan(&appAidName)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshAid()-Raw", tx.Error.Error()))
	}
	// 查询结果为空则删除容器元素
	if tx.RowsAffected == 0 {
		var checkApp = entity.App{Model: gorm.Model{ID: appId}}
		if err := global.MDB.First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("appId为%d的应用不存在或者无效！", appId))
		}
		delete(AidByAppMap, checkApp.Host)
		return nil
	}
	// 刷新容器
	AidByAppMap[appAidName.Host] = strings.Split(appAidName.AidName, config.Program.AidDelimiter)
	global.AppLogger.Info(fmt.Sprintf("complete the initialization of Aid container under the application with appId %d", appId))
	return nil
}
