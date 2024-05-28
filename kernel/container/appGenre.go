package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/global"
	"loiter/model/entity"
)

/**
 * 应用类型容器
 * @auth eyesYeager
 * @date 2024/3/6 16:38
 */

// GenreByAppMap 应用类型 by AppHost
var GenreByAppMap = make(map[string]string)

// InitAppGenre 初始化应用类型容器
func InitAppGenre() {
	global.AppLogger.Info("start initializing the AppGenre container")
	// 获取有效应用
	var appList []entity.App
	if err := global.MDB.Where(&entity.App{Status: constant.Status.Normal.Code}).Find(&appList).Error; err != nil {
		panic(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if appList == nil || len(appList) == 0 {
		global.AppLogger.Warn("there are currently no valid app requiring service")
		return
	}
	// 构建容器
	containerMap := make(map[string]string)
	for _, item := range appList {
		containerMap[item.Host] = item.Genre
	}
	GenreByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of AppGenre container")
}

// RefreshAppGenre 刷新应用类型容器
func RefreshAppGenre(appId uint) error {
	// 获取有效应用
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新应用类型容器失败，应用不存在！", appId))
	}
	// 若状态为无效，则删除
	if checkApp.Status != constant.Status.Normal.Code {
		delete(GenreByAppMap, checkApp.Host)
		return nil
	}
	// 刷新容器
	GenreByAppMap[checkApp.Host] = checkApp.Genre
	return nil
}

// DeleteAppGenre 删除应用类型容器
func DeleteAppGenre(host string) {
	delete(GenreByAppMap, host)
}
