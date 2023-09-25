package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"loiter/config"
	"loiter/constant"
	"loiter/global"
	"os"
	"path/filepath"
	"time"
)

/**
 * 日志工具初始化
 * 参考链接：https://www.bilibili.com/video/BV1L24y1q7DA
 * @author eyesYeager
 * @date 2023/7/2 16:31
 */

// logBootstrap 日志初始化入口方法
func logBootstrap() {
	global.AppLogger = initAppLogger()
	global.BackstageLogger = initBackstageLogger()
	global.GatewayLogger = initGatewayLogger()
}

func initAppLogger() *zap.SugaredLogger {
	path := config.Program.LogBasePath + string(filepath.Separator) + config.Program.LogAppPath
	return initLogger(path)
}

func initBackstageLogger() *zap.SugaredLogger {
	path := config.Program.LogBasePath + string(filepath.Separator) + config.Program.LogBackstagePath
	return initLogger(path)
}

func initGatewayLogger() *zap.SugaredLogger {
	path := config.Program.LogBasePath + string(filepath.Separator) + config.Program.LogGateWayPath
	return initLogger(path)
}

// initLogger 初始化日志对象
func initLogger(path string) *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(path), getLogLevel())
	return zap.New(core).Sugar()
}

// getEncoder 配置日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 时间格式化
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	// JSON格式打印
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getWriteSyncer 配置文件打印方式
func getWriteSyncer(path string) zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + path + stSeparator + time.Now().Format(time.DateOnly) + "." + config.Program.LogSuffix

	// 单个文件超出大小时进行文件分割
	lumberjackSyncer := &lumberjack.Logger{
		Filename: stLogFilePath,
		MaxSize:  config.Program.LogMaxSize,  // 文件最大尺寸(MB)
		MaxAge:   config.Program.LogMaxAge,   // 文件保存最长时间
		Compress: config.Program.LogCompress, // 是否压缩
	}

	// 开发模式下，日志会输出到控制台，但生产模式不会
	if config.Program.ProgramMode == constant.DEVELOP {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackSyncer), zapcore.AddSync(os.Stdout))
	} else {
		return zapcore.AddSync(lumberjackSyncer)
	}

}

// getLogLevel 设置日志级别，开发模式下为debug级别，生产模式为Info级别
func getLogLevel() zapcore.Level {
	var logMode zapcore.Level
	if config.Program.ProgramMode == constant.DEVELOP {
		logMode = zapcore.DebugLevel
	} else {
		logMode = zapcore.InfoLevel
	}
	return logMode
}
