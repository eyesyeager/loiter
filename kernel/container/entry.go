package container

import (
	"fmt"
	"loiter/global"
)

/**
 * 容器总管理
 * @auth eyesYeager
 * @date 2024/1/12 11:19
 */

// InitRegister 初始化所有注册信息
func InitRegister() {
	global.BackstageLogger.Info("start initializing all containers")
	InitAppServer()
	InitBalancer()
	InitPassageway()
	InitLimiter()
	InitNameList()
	InitAid()
	global.BackstageLogger.Info("all containers initialization completed")
}

// RefreshRegister 刷新所有注册信息
func RefreshRegister(appId uint) error {
	global.BackstageLogger.Info(fmt.Sprintf("start refreshing all containers under the application with appId %d", appId))
	if err := RefreshAppServer(appId); err != nil {
		return err
	}
	if err := RefreshBalancer(appId); err != nil {
		return err
	}
	if err := RefreshPassageway(appId); err != nil {
		return err
	}
	if err := RefreshLimiter(appId); err != nil {
		return err
	}
	if err := RefreshNameList(appId); err != nil {
		return err
	}
	if err := RefreshAid(appId); err != nil {
		return err
	}
	global.BackstageLogger.Info(fmt.Sprintf("complete the refresh of all containers under the application with appId %d", appId))
	return nil
}
