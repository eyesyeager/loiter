package container

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/model/po"
	"loiter/plugin/passageway/filter/limiter"
)

/**
 * 限流容器
 * @auth eyesYeager
 * @date 2024/1/12 15:47
 */

// LimiterByAppMap 限流器 by AppHost
var LimiterByAppMap map[string]limiter.ILimiter

// InitLimiter 初始化限流容器
func InitLimiter() {
	global.AppLogger.Info("start initializing the Limiter container")
	// 获取限流配置
	var appLimiterName []po.GetAppLimiterName
	if rowsAffected := global.MDB.Raw(`SELECT a.host, al.limiter_name, al.parameter
						FROM app a, app_limiter al
						WHERE a.status = ? AND a.id = al.app_id`, constant.Status.Normal).Scan(&appLimiterName).RowsAffected; rowsAffected == 0 {
		global.AppLogger.Info("there is currently no valid limiter configuration")
		return
	}
	// 构建限流容器
	var containerMap = make(map[string]limiter.ILimiter)
	for _, item := range appLimiterName {
		err, iLimiter := limiter.NewLimiterFilter(item.LimiterName, item.Parameter)
		if err != nil {
			global.AppLogger.Error(fmt.Sprintf("failed to build the Limiter container for application with host %s, error message: %s", item.Host, err.Error()))
			continue
		}
		containerMap[item.Host] = iLimiter
	}
	LimiterByAppMap = containerMap
	global.AppLogger.Info("complete the initialization of Limiter container")
}

// RefreshLimiter 刷新限流容器
func RefreshLimiter(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the Limiter container under the application with appId %d", appId))
	// 获取限流配置
	var appLimiterName po.GetAppLimiterName
	tx := global.MDB.Raw(`SELECT a.host, al.limiter_name, al.parameter
				FROM app a, app_limiter al
				WHERE a.id = ? AND a.status = ? AND a.id = al.app_id`, appId, constant.Status.Normal).Scan(&appLimiterName)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshLimiter()-Raw", tx.Error.Error()))
	}
	// 如果想关闭限流，应该去改passageway，而不是让limiter为空，因此如果配置为空，那就直接返回
	if tx.RowsAffected == 0 {
		return nil
	}
	// 刷新限流容器
	err, iLimiter := limiter.NewLimiterFilter(appLimiterName.LimiterName, appLimiterName.Parameter)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to refresh the Limiter container for application with appId %d, error message: %s", appId, err.Error()))
	}
	LimiterByAppMap[appLimiterName.Host] = iLimiter
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of Limiter container under the application with appId %d", appId))
	return nil
}
