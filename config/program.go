package config

import (
	"loiter/constant"
	"path/filepath"
)

/**
 * @author eyesYeager
 * @date 2023/7/2 13:19
 */

// Program 程序配置
var Program = programConfig{
	Mode:                         constant.DEVELOP,    // 程序运行模式
	Name:                         "Loiter",            // 应用名称
	GateWayPort:                  "9500",              // 网关服务端口号
	BackstagePort:                "9510",              // 后台服务端口号
	LogConfig:                    logConfigInstance,   // 日志配置
	MySQLConfig:                  mysqlConfigInstance, // MySQL配置
	AESSecretKey:                 "hello,bestLoiter",  // AES双向加密密钥，必须是16位
	JWTSecretKey:                 "hello,loiter!",     // JWT签名加密密钥
	JWTExpire:                    500,                 // JWT签名过期时间(min)
	EmailConfig:                  emailConfigInstance, // 邮箱配置
	InitialPsdLen:                8,                   // 初始密码长度
	ServerDefaultNameLen:         10,                  // 默认应用实例名长度
	BalancerDefaultStrategy:      1,                   // 默认负载均衡策略(值为策略id)
	DefaultPassageway:            "requestLog",        // 通道插件配置-默认通道配置
	PassagewayDelimiter:          ",",                 // 通道插件配置-分隔符
	DefaultAid:                   "requestFill",       // 响应处理器插件配置-默认响应处理器配置
	AidDelimiter:                 ",",                 // 响应处理器插件配置-分隔符
	NameListBloomCapacity:        10000,               // 黑白名单插件配置-布隆过滤器参数-容量
	NameListBloomMisjudgmentRate: 0.01,                // 黑白名单插件配置-布隆过滤器参数-误判率
	NameListIpDelimiter:          ";",                 // 黑白名单插件配置-后台接收ip操作相关请求使用的分隔符
}

// logConfigInstance 日志配置
var logConfigInstance = logConfig{
	LogSuffix:        "log",                                          // 日志文件拓展名
	LogBasePath:      "runtime" + string(filepath.Separator) + "log", // 日志文件基础路径
	LogAppPath:       "app",                                          // 系统日志文件路径
	LogBackstagePath: "backstage",                                    // 管理系统日志文件路径
	LogGateWayPath:   "gateway",                                      // 网关日志文件路径
	LogSQLPath:       "sql",                                          // SQL日志文件路径
	LogMaxSize:       50,                                             // 单个日志文件最大尺寸(MB)
	LogMaxAge:        30,                                             // 单个日志文件最多保存多少天
	LogCompress:      true,                                           // 日志文件是否开启压缩
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
	Password: "",                   // 密码
	Host:     "smtp.163.com",       // 主机地址
}

type programConfig struct {
	// Develop
	Mode int // 程序运行模式

	// Application
	Name          string // 应用名称
	GateWayPort   string // 网关服务端口号
	BackstagePort string // 后台服务端口号

	// Log
	LogConfig logConfig

	// MySQL
	MySQLConfig mysqlConfig

	// secret key
	AESSecretKey string // AES双向加密密钥，必须是16位

	// JWT
	JWTSecretKey string // JWT签名加密密钥
	JWTExpire    int    // JWT签名过期时间(min)

	// email
	EmailConfig emailConfig

	// User
	InitialPsdLen int // 初始密码长度

	// 默认应用实例名长度
	ServerDefaultNameLen int

	// 默认负载均衡策略(值为策略id)
	BalancerDefaultStrategy uint

	// 默认通道配置
	DefaultPassageway string

	// 通道配置分隔符
	PassagewayDelimiter string

	// 默认响应处理器配置
	DefaultAid string

	// 响应处理器配置分割符
	AidDelimiter string

	// 黑白名单插件配置
	NameListBloomCapacity        uint    // 布隆过滤器参数-容量
	NameListBloomMisjudgmentRate float64 // 布隆过滤器参数-误判率
	NameListIpDelimiter          string  // 后台接收ip操作相关请求使用的分隔符
}

type logConfig struct {
	LogSuffix        string // 日志文件拓展名
	LogBasePath      string // 日志文件基础路径
	LogAppPath       string // 系统日志文件路径
	LogBackstagePath string // 管理系统日志文件路径
	LogGateWayPath   string // 网关日志文件路径
	LogSQLPath       string // SQL日志文件路径
	LogMaxSize       int    // 单个日志文件最大尺寸(MB)
	LogMaxAge        int    // 单个日志文件最多保存多少天
	LogCompress      bool   // 日志文件是否开启压缩
}

type mysqlConfig struct {
	Host                string // 主机地址
	Port                int    // 端口号
	Database            string // 数据库名
	Username            string // 用户名
	Password            string // 用户密码
	Charset             string // 编码格式
	MaxIdleConn         int    // 空闲连接池中连接的最大数量
	MaxOpenConn         int    // 打开数据库连接的最大数量
	LogMode             string // 日志级别
	EnableFileLogWriter bool   // 是否启用日志文件
}

type emailConfig struct {
	Addr     string // SMTP服务器的地址
	Identity string // 身份证明
	Username string // 用户名
	Password string // 密码
	Host     string // 主机地址
}
