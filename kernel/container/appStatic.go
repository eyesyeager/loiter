package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
)

/**
 * 应用静态配置容器
 * @author eyesYeager
 * @date 2024/3/21 10:02
 */

var StaticByAppMap map[string]AppStatic

type AppStatic struct {
	ErrorRoute string
}

// InitAppStatic 初始化应用静态配置容器
func InitAppStatic() {
	global.AppLogger.Info("start initializing the AppStatic container")
	var appList []entity.App
	if tx := global.MDB.Where(&entity.App{
		Status: constant.Status.Normal.Code,
		Genre:  constants.AppGenre.Static,
	}).Find(&appList); tx.RowsAffected == 0 {
		global.AppLogger.Warn("there are currently no valid static app")
		return
	}
	containerMap := make(map[string]AppStatic)
	for _, item := range appList {
		containerMap[item.Host] = AppStatic{
			ErrorRoute: item.ErrorRoute,
		}
	}
	StaticByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of AppStatic container")
}

// RefreshAppStatic 刷新指定应用静态配置容器
func RefreshAppStatic(appId uint) error {
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新静态配置容器失败，应用不存在！", appId))
	}
	// 若类型不为static，或者状态不为normal，则删除容器项
	if checkApp.Genre != constants.AppGenre.Static || checkApp.Status != constant.Status.Normal.Code {
		DeleteAppStatic(checkApp.Host)
		return nil
	}
	// 刷新容器
	StaticByAppMap[checkApp.Host] = AppStatic{
		ErrorRoute: checkApp.ErrorRoute,
	}
	return nil
}

// DeleteAppStatic 删除指定应用静态配置容器
func DeleteAppStatic(host string) {
	delete(StaticByAppMap, host)
}
