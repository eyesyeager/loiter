package config

import (
	"loiter/constant"
	"path/filepath"
)

/**
 * @author eyesYeager
 * @date 2023/7/2 13:19
 */

var Program = programConfig{
	constant.DEVELOP,
	"loiter",
	"9500",
	"9510",
	"log",
	"runtime" + string(filepath.Separator) + "log",
	"app",
	"backstage",
	"gateway",
	"sql",
	50,
	30,
	true,
	mysqlConfigInstance,
	"hello,bestLoiter",
	"hello,loiter!",
	5,
	emailConfigInstance,
	10,
}

var mysqlConfigInstance = mysqlConfig{
	"192.168.204.133",
	3306,
	"loiter",
	"root",
	"root",
	"utf8mb4",
	5,
	10,
	"info",
	true,
}

var emailConfigInstance = emailConfig{
	"",
	"",
	"",
	"",
	"",
}

type programConfig struct {
	// Develop
	ProgramMode int // 程序运行模式

	// Application
	Name          string // 应用名称
	GateWayPort   string // 网关服务端口号
	BackstagePort string // 后台服务端口号

	// Log
	LogSuffix        string // 日志文件拓展名
	LogBasePath      string // 日志文件基础路径
	LogAppPath       string // 系统日志文件路径
	LogBackstagePath string // 管理系统日志文件路径
	LogGateWayPath   string // 网关日志文件路径
	LogSQLPath       string // SQL日志文件路径
	LogMaxSize       int    // 单个日志文件最大尺寸(MB)
	LogMaxAge        int    // 单个日志文件最多保存多少天
	LogCompress      bool   // 日志文件是否开启压缩

	// MySQL
	MySQLConfig mysqlConfig

	// secret key
	AESSecretKey string // AES双向加密密钥，必须是16位

	// JWT
	JWTSecretKey string // JWT签名加密密钥
	JWTExpire    int    // JWT签名过期时间(min)

	// email
	EmailConfig emailConfig

	// 默认应用实例名长度
	ServerDefaultNameLen int
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
