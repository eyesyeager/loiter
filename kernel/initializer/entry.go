package initializer

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"loiter/app/plugin/aid"
	"loiter/app/plugin/balancer"
	"loiter/app/plugin/exception"
	"loiter/app/plugin/filter"
	"loiter/app/plugin/filter/limiter"
	"loiter/app/plugin/final"
	"loiter/config"
	"loiter/global"
	"loiter/model/entity"
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
	initProcessorData()
	initLimiterData()
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
	var roleEntityList []entity.Role
	var valInfo = reflect.ValueOf(config.RoleConfig)
	num := valInfo.NumField()
	if num <= 0 {
		panic("role information cannot be empty, please check config.RoleConfig")
	}
	for i := 0; i < num; i++ {
		val := valInfo.Field(i).Interface()
		roleEntityList = append(roleEntityList, val.(config.RoleStruct).Entity)
	}
	if err := global.MDB.Create(&roleEntityList).Error; err != nil {
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
	for index := range config.UserConfig {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(config.UserConfig[index].Password), bcrypt.DefaultCost)
		if err != nil {
			panic(fmt.Errorf("user data initialization failed! An error occurred while handling initial password: %s", err.Error()))
		}
		config.UserConfig[index].Password = string(passwordHash)
	}
	if err := global.MDB.Create(&config.UserConfig).Error; err != nil {
		panic(fmt.Errorf("user data initialization failed! An error occurred while creating the user data: %s", err.Error()))
	}
	global.AppLogger.Info("user data initialization completed")
}

// initBalancerData 初始化负载均衡数据
func initBalancerData() {
	global.AppLogger.Info("start initializing Balancer data")
	// 不允许没有可用的负载均衡插件
	if len(balancer.IBalancerConfigList) == 0 {
		panic(errors.New("not allowed No Balancer plugin available"))
	}
	// 清空所有负载均衡器数据
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Balancer{}).Error; err != nil {
		panic(fmt.Errorf("balancer data initialization failed! An error occurred while clearing all Balancer data: %s", err.Error()))
	}
	// 插入新数据
	if err := global.MDB.Create(&balancer.IBalancerConfigList).Error; err != nil {
		panic(fmt.Errorf("balancer data initialization failed! An error occurred while creating Balancer data: %s", err.Error()))
	}
	global.AppLogger.Info("balancer data initialization completed")
}

// initProcessorData 初始化处理器数据
func initProcessorData() {
	global.AppLogger.Info("start initializing processor data")
	// 清空所有处理器
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Processor{}).Error; err != nil {
		panic(fmt.Errorf("processor data initialization failed! An error occurred while clearing all processor data: %s", err.Error()))
	}
	// 插入新数据
	var processorList []entity.Processor
	processorList = append(processorList, filter.IFilterConfigList...)
	processorList = append(processorList, aid.IAidConfigList...)
	processorList = append(processorList, exception.IExceptionConfigList...)
	processorList = append(processorList, final.IFinalConfigList...)
	if len(processorList) != 0 {
		if err := global.MDB.Create(&processorList).Error; err != nil {
			panic(fmt.Errorf("processor data initialization failed! An error occurred while creating processor data: %s", err.Error()))
		}
	}
	global.AppLogger.Info("processor data initialization completed")
}

// initLimiterData 初始化限流器数据
func initLimiterData() {
	global.AppLogger.Info("start initializing limiter data")
	// 清空所有限流器
	if err := global.MDB.Where("1 = 1").Unscoped().Delete(&entity.Limiter{}).Error; err != nil {
		panic(fmt.Errorf("limiter data initialization failed! An error occurred while clearing all limiter data: %s", err.Error()))
	}
	// 插入新数据
	var limiterEntityList []entity.Limiter
	var valInfo = reflect.ValueOf(limiter.LimiterConfig)
	num := valInfo.NumField()
	if num != 0 {
		for i := 0; i < num; i++ {
			val := valInfo.Field(i).Interface()
			limiterEntityList = append(limiterEntityList, val.(entity.Limiter))
		}
		if err := global.MDB.Create(&limiterEntityList).Error; err != nil {
			panic(fmt.Errorf("limiter data initialization failed! An error occurred while creating limiter data: %s", err.Error()))
		}
	}
	global.AppLogger.Info("limiter data initialization completed")
}
