package initializer

import (
	"fmt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/model/entity"
	"loiter/plugin/aid"
	"loiter/plugin/loadbalancer"
	"loiter/plugin/passageway"
	"loiter/plugin/passageway/filter/limiter"
	"reflect"
)

/**
 * 初始数据加载器
 * @auth eyesYeager
 * @date 2024/1/11 15:01
 */

// InitData 系统数据初始化
func InitData() {
	global.AppLogger.Info("start initializing system data")
	initRoleData()
	initUserData()
	initBalancerData()
	initLimiterData()
	initPassagewayData()
	initAidData()
	global.AppLogger.Info("system data initialization completed")
}

// initRoleData 初始化角色数据
func initRoleData() {
	global.AppLogger.Info("start initializing role data")
	var count int64
	if err := global.MDB.Model(&entity.Role{}).Count(&count).Error; err != nil {
		panic(fmt.Errorf("role data initialization failed! An error occurred while querying the total number of characters: %s", err.Error()))
	}
	// 如果现有角色不为空，则略过
	if count != 0 {
		global.AppLogger.Info("the role data already exists and no initialization is required")
		return
	}
	// 如果现有角色为空，则初始化
	var roleEntitySlice []entity.Role
	var valInfo = reflect.ValueOf(config.RoleConfig)
	num := valInfo.NumField()
	for i := 0; i < num; i++ {
		val := valInfo.Field(i).Interface()
		roleEntitySlice = append(roleEntitySlice, val.(config.RoleStruct).Entity)
	}
	if err := global.MDB.Create(&roleEntitySlice).Error; err != nil {
		panic(fmt.Errorf("role data initialization failed! An error occurred while creating the role data: %s", err.Error()))
	}
	global.AppLogger.Info("role data initialization completed")
}

// initUserData 初始化用户数据
func initUserData() {
	global.AppLogger.Info("start initializing user data")
	var count int64
	if err := global.MDB.Model(&entity.User{}).Count(&count).Error; err != nil {
		panic(fmt.Errorf("user data initialization failed! An error occurred while querying the total number of characters: %s", err.Error()))
	}
	// 如果现有用户不为空，则略过
	if count != 0 {
		global.AppLogger.Info("the user data already exists and no initialization is required")
		return
	}
	// 如果现有用户为空，则初始化
	var userEntitySlice []entity.User
	var valInfo = reflect.ValueOf(config.UserConfig)
	num := valInfo.NumField()
	for i := 0; i < num; i++ {
		val := valInfo.Field(i).Interface()
		userEntitySlice = append(userEntitySlice, val.(entity.User))
	}
	if err := global.MDB.Create(&userEntitySlice).Error; err != nil {
		panic(fmt.Errorf("user data initialization failed! An error occurred while creating the user data: %s", err.Error()))
	}
	global.AppLogger.Info("user data initialization completed")
}

// initBalancerData 初始化负载均衡数据
func initBalancerData() {
	global.AppLogger.Info("start initializing Balancer data")
	// 清空所有负载均衡器数据
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Balancer{}).Error; err != nil {
		panic(fmt.Errorf("balancer data initialization failed! An error occurred while clearing all Balancer data: %s", err.Error()))
	}
	// 插入新数据
	if err := global.MDB.Create(&loadbalancer.BalancerConfigSlice).Error; err != nil {
		panic(fmt.Errorf("balancer data initialization failed! An error occurred while creating Balancer data: %s", err.Error()))
	}
	global.AppLogger.Info("balancer data initialization completed")
}

// initLimiterData 初始化限流器数据
func initLimiterData() {
	global.AppLogger.Info("start initializing limiter data")
	// 清空所有限流器
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Limiter{}).Error; err != nil {
		panic(fmt.Errorf("limiter data initialization failed! An error occurred while clearing all limiter data: %s", err.Error()))
	}
	// 插入新数据
	var limiterEntitySlice []entity.Limiter
	var valInfo = reflect.ValueOf(limiter.LimiterConfig)
	num := valInfo.NumField()
	for i := 0; i < num; i++ {
		val := valInfo.Field(i).Interface()
		limiterEntitySlice = append(limiterEntitySlice, val.(entity.Limiter))
	}
	if err := global.MDB.Create(&limiterEntitySlice).Error; err != nil {
		panic(fmt.Errorf("limiter data initialization failed! An error occurred while creating limiter data: %s", err.Error()))
	}
	global.AppLogger.Info("limiter data initialization completed")
}

// initPassagewayData 初始化通道数据
func initPassagewayData() {
	global.AppLogger.Info("start initializing passageway data")
	// 清空所有通道
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Passageway{}).Error; err != nil {
		panic(fmt.Errorf("passageway data initialization failed! An error occurred while clearing all passageway data: %s", err.Error()))
	}
	// 插入新数据
	if err := global.MDB.Create(&passageway.PassageConfigSlice).Error; err != nil {
		panic(fmt.Errorf("passageway data initialization failed! An error occurred while creating passageway data: %s", err.Error()))
	}
	global.AppLogger.Info("passageway data initialization completed")
}

// initAidData 初始化响应处理器数据
func initAidData() {
	global.AppLogger.Info("start initializing aid data")
	// 清空所有响应处理器
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Aid{}).Error; err != nil {
		panic(fmt.Errorf("aid data initialization failed! An error occurred while clearing all aid data: %s", err.Error()))
	}
	// 插入新数据
	if err := global.MDB.Create(&aid.IAidConfigSlice).Error; err != nil {
		panic(fmt.Errorf("aid data initialization failed! An error occurred while creating aid data: %s", err.Error()))
	}
	global.AppLogger.Info("aid data initialization completed")
}
