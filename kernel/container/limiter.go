package container

import (
	"fmt"
	"loiter/global"
)

/**
 * 限流容器
 * @auth eyesYeager
 * @date 2024/1/12 15:47
 */

// LimiterByAppMap 限流器 by AppHost
var LimiterByAppMap map[string]string

// InitLimiter 初始化限流容器
func InitLimiter() {
	global.AppLogger.Info("start initializing the Limiter container")
	//var appLimiterName []po.GetAppLimiterName
	//if rowsAffected := global.MDB.Raw(`SELECT a.host, al.name limiterName, al.parameter
	//			FROM app a, app_limiter al
	//			WHERE a.status = ? AND a.id = al.app_id`, constant.Status.Normal).Scan(&appLimiterName).RowsAffected; rowsAffected == 0 {
	//	global.AppLogger.Warn("there is currently no valid limiter configuration")
	//}

	global.AppLogger.Info("complete the initialization of Limiter container")
}

// RefreshLimiter 刷新限流容器
func RefreshLimiter(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Limiter container under the application with appId %d", appId))

	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Limiter container under the application with appId %d", appId))
	return nil
}
