package bootstrap

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/model/entity"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

/**
 * MySQL启动文件
 * @author eyesYeager
 * @date 2023/4/10 13:42
 */

// 初始化 MySQL
func mDbBootstrap() {
	dbConfig := config.Program.MySQLConfig
	if dbConfig.Database == "" {
		panic(fmt.Errorf("mysql database name cannot be empty"))
	}
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}); err != nil {
		panic(fmt.Errorf("mysql connect failed: %s", err))
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
		initMySqlTables(db)
		global.MDB = db
	}
}

// mDbClose 关闭 MySQL 连接
func mDbClose() {
	if global.MDB != nil {
		db, _ := global.MDB.DB()
		_ = db.Close()
	}
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.App{},
		entity.Role{},
		entity.User{},
		entity.LogLogin{},
	)
	if err != nil {
		global.AppLogger.Error("migrate table failed")
		os.Exit(0)
	}
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if config.Program.MySQLConfig.EnableFileLogWriter {
		stSeparator := string(filepath.Separator)
		stRootDir, _ := os.Getwd()
		path := config.Program.LogBasePath + stSeparator + config.Program.LogSQLPath
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename: stRootDir + stSeparator + path + stSeparator + time.Now().Format(time.DateOnly) + "." + config.Program.LogSuffix,
			MaxSize:  config.Program.LogMaxSize,
			MaxAge:   config.Program.LogMaxAge,
			Compress: config.Program.LogCompress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch config.Program.MySQLConfig.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !config.Program.MySQLConfig.EnableFileLogWriter, // 禁用彩色打印
	})
}
