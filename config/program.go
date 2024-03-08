package config

import (
	"loiter/constants"
	"path/filepath"
)

/**
 * @author eyesYeager
 * @date 2023/7/2 13:19
 */

// Program 程序配置
var Program = programConfig{
	Mode:                 constants.DEVELOP,    // 程序运行模式
	Name:                 "Loiter",             // 应用名称
	GateWayPort:          "9500",               // 网关服务端口号
	BackstagePort:        "9510",               // 后台服务端口号
	StaticDirPath:        "/static",            // 静态资源存放路径
	LogConfig:            logConfigInstance,    // 日志配置
	MySQLConfig:          mysqlConfigInstance,  // MySQL配置
	AESSecretKey:         "hello,bestLoiter",   // AES双向加密密钥，必须是16位
	JWTSecretKey:         "hello,loiter!",      // JWT签名加密密钥
	JWTExpire:            1000,                 // JWT签名过期时间(min)
	EmailConfig:          emailConfigInstance,  // 邮箱配置
	InitialPsdLen:        8,                    // 初始密码长度
	ServerDefaultNameLen: 10,                   // 默认应用实例名长度
	PluginConfig:         pluginConfigInstance, // 插件配置
}

// logConfigInstance 日志配置
var logConfigInstance = logConfig{
	LogSuffix:      "log",                                          // 日志文件拓展名
	LogBasePath:    "runtime" + string(filepath.Separator) + "log", // 日志文件基础路径
	LogAppPath:     "app",                                          // 系统日志文件路径
	LogGateWayPath: "gateway",                                      // 网关日志文件路径
	LogSQLPath:     "sql",                                          // SQL日志文件路径
	LogMaxSize:     50,                                             // 单个日志文件最大尺寸(MB)
	LogMaxAge:      30,                                             // 单个日志文件最多保存多少天
	LogCompress:    true,                                           // 日志文件是否开启压缩
}

// mysqlConfigInstance MySQL配置
var mysqlConfigInstance = mysqlConfig{
	Host:                "127.0.0.1", // 主机地址
	Port:                3306,        // 端口号
	Database:            "loiter",    // 数据库名
	Username:            "root",      // 用户名
	Password:            "792734338", // 用户密码
	Charset:             "utf8mb4",   // 编码格式
	MaxIdleConn:         5,           // 空闲连接池中连接的最大数量
	MaxOpenConn:         10,          // 打开数据库连接的最大数量
	LogMode:             "info",      // 日志级别
	EnableFileLogWriter: true,        // 是否启用日志文件
}

// emailConfigInstance 邮箱配置
var emailConfigInstance = emailConfig{
	Addr:     "smtp.163.com:25",    // SMTP服务器的地址
	Identity: "",                   // 身份证明
	Username: "eyesyeager@163.com", // 用户名
	Password: "VAUARZWCFWSXGEUP",   // 密码
	Host:     "smtp.163.com",       // 主机地址
}

// pluginConfigInstance 插件配置
var pluginConfigInstance = pluginConfig{
	BalancerDefaultStrategy:      "random",     // 默认插件-负载均衡
	FilterDefaultStrategy:        "",           // 默认插件-过滤器(多个则用 ProcessorDelimiter 隔开)
	AidDefaultStrategy:           "",           // 默认插件-响应处理器(多个则用 ProcessorDelimiter 隔开)
	ExceptionDefaultStrategy:     "notice",     // 默认插件-异常处理器(多个则用 ProcessorDelimiter 隔开)
	FinalDefaultStrategy:         "requestLog", // 默认插件-最终处理器(多个则用 ProcessorDelimiter 隔开)
	ProcessorDelimiter:           ",",          // 处理器-分隔符
	NameListBloomCapacity:        10000,        // 过滤器-黑白名单插件配置-布隆过滤器参数-容量
	NameListBloomMisjudgmentRate: 0.01,         // 过滤器-黑白名单插件配置-布隆过滤器参数-误判率
	NameListIpDelimiter:          ";",          // 过滤器-黑白名单插件配置-后台接收ip操作相关请求使用的分隔符
}

type programConfig struct {
	// Develop
	Mode int

	// Application
	Name          string
	GateWayPort   string
	BackstagePort string
	StaticDirPath string

	// Log
	LogConfig logConfig

	// MySQL
	MySQLConfig mysqlConfig

	// secret key
	AESSecretKey string

	// JWT
	JWTSecretKey string
	JWTExpire    int

	// email
	EmailConfig emailConfig

	// User
	InitialPsdLen int

	// 默认应用实例名长度
	ServerDefaultNameLen int

	// plugin
	PluginConfig pluginConfig
}

type logConfig struct {
	LogSuffix      string
	LogBasePath    string
	LogAppPath     string
	LogGateWayPath string
	LogSQLPath     string
	LogMaxSize     int
	LogMaxAge      int
	LogCompress    bool
}

type mysqlConfig struct {
	Host                string
	Port                int
	Database            string
	Username            string
	Password            string
	Charset             string
	MaxIdleConn         int
	MaxOpenConn         int
	LogMode             string
	EnableFileLogWriter bool
}

type emailConfig struct {
	Addr     string
	Identity string
	Username string
	Password string
	Host     string
}

type pluginConfig struct {
	BalancerDefaultStrategy      string
	FilterDefaultStrategy        string
	AidDefaultStrategy           string
	ExceptionDefaultStrategy     string
	FinalDefaultStrategy         string
	ProcessorDelimiter           string
	NameListBloomCapacity        uint
	NameListBloomMisjudgmentRate float64
	NameListIpDelimiter          string
}
