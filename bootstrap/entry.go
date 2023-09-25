package bootstrap

import "loiter/global"

/**
 * @author eyesYeager
 * @date 2023/7/2 16:31
 */

func Start() {
	logBootstrap()
	global.AppLogger.Info("log tool initialization completed")
	mDbBootstrap()
	global.AppLogger.Info("mysql tool initialization completed")
}
