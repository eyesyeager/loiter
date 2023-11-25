package bootstrap

import "loiter/global"

/**
 * @author eyesYeager
 * @date 2023/7/2 16:31
 */

// Start 程序环境启动入口
func Start() {
	logBootstrap()
	global.AppLogger.Info("log tool initialization completed")
	mDbBootstrap()
	global.AppLogger.Info("mysql tool initialization completed")
}

// End 程序环境关闭入口
func End() {
	mDbClose()
	global.AppLogger.Info("MySQL connection closed")
}
