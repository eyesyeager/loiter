package initializer

import (
	"loiter/global"
	"loiter/kernel/model/entity"
)

/**
 * 初始数据加载器
 * 第一版先不做，以后再说
 * @auth eyesYeager
 * @date 2024/1/11 15:01
 */

// InitData 系统数据初始化
func InitData() {
	global.AppLogger.Info("start initializing system data")
	initRoleData()
	initUserData()
	initBalancerData()
	initPassagewayData()
	global.AppLogger.Info("system data initialization completed")
}

// initRoleData 初始化角色数据
func initRoleData() {
	global.AppLogger.Info("start initializing system data")
	// 如果现有角色不为空，则掠过
	global.MDB.Model(&entity.Role{})
	// 如果现有角色为空，则初始化
	global.AppLogger.Info("system data initialization completed")
}

// initUserData 初始化用户数据
func initUserData() {
	global.AppLogger.Info("start initializing system data")
	// 如果现有用户不为空，则掠过
	// 如果现有用户为空，则初始化
	global.AppLogger.Info("system data initialization completed")
}

// initBalancerData 初始化负载均衡数据
func initBalancerData() {
	global.AppLogger.Info("start initializing system data")
	// 清空所有负载均衡器数据
	// 获取用户自定义均衡器
	// 组装数据
	// 插入新数据
	global.AppLogger.Info("system data initialization completed")
}

// initPassagewayData 初始化通道数据
func initPassagewayData() {
	global.AppLogger.Info("start initializing system data")
	// 清空所有通道
	// 获取用户自定义通道
	// 组装数据
	// 插入新数据
	global.AppLogger.Info("system data initialization completed")
}
