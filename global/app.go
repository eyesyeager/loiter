package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	config2 "zliway/bootstrap/config"
)

/**
 * 全局变量
 * @author eyesYeager
 * @date 2023/4/10 16:00
 */

// Profiles 启动环境
var Profiles = new(config2.Profiles)

// 当前环境配置集合
type configuration struct {
	App        config2.App
	Log        config2.Log
	Persistent config2.Persistent
	Program    config2.Program
}

// Config 配置文件
var Config = new(configuration)

// Log 日志
var Log = new(zap.Logger)

// MDB mysql
var MDB = new(gorm.DB)

// RDB redis
var RDB = new(redis.Client)
