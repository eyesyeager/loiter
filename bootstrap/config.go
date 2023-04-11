package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"zliway/global"
)

/**
 * 配置启动文件
 * @author eyesYeager
 * @date 2023/4/10 13:42
 */

// 默认文件名，可以根据自身情况修改
var fileName = "config"

// 默认文件后缀，不建议修改
var fileSuffix = ".yaml"

// InitializeConfig 初始化配置信息
func InitializeConfig() {
	// 初始化变量
	v := viper.New()
	rootFile := fileName + fileSuffix

	// 读取根配置文件
	v.SetConfigFile(rootFile)
	v.SetConfigType(fileSuffix[1:])
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("configuration file '%s' read failed: %s", rootFile, err))
	}

	// 初始化全局变量global.Profiles
	if err := v.Unmarshal(&global.Profiles); err != nil {
		panic(fmt.Errorf("configuration file '%s' assignment failed: %s", rootFile, err))
	}

	// 读取当前环境下的配置文件
	currentFile := fileName + "-" + global.Profiles.Active + fileSuffix
	v.SetConfigFile(currentFile)
	v.SetConfigType(fileSuffix[1:])
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("configuration file '%s' read failed: %s", rootFile, err))
	}

	// 初始化全局变量global.Config
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("configuration file '%s' assignment failed: %s", rootFile, err))
	}
}
