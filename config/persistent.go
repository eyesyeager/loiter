package config

/**
 * 持久层配置
 * @author eyesYeager
 * @date 2023/4/10 17:23
 */

type Persistent struct {
	Mysql Mysql
	Redis Redis
}

// Mysql MySQL配置
type Mysql struct {
	Host                string `mapstructure:"host" yaml:"host"`                               // 主机地址
	Port                int    `mapstructure:"port" yaml:"port"`                               // 端口号
	Database            string `mapstructure:"database" yaml:"database"`                       // 数据库名
	Username            string `mapstructure:"username" yaml:"username"`                       // 用户名
	Password            string `mapstructure:"password" yaml:"password"`                       // 用户密码
	Charset             string `mapstructure:"charset" yaml:"charset"`                         // 编码格式
	MaxIdleConn         int    `mapstructure:"maxIdleConn" yaml:"maxIdleConn"`                 // 空闲连接池中连接的最大数量
	MaxOpenConn         int    `mapstructure:"maxOpenConn" yaml:"maxOpenConn"`                 // 打开数据库连接的最大数量
	LogMode             string `mapstructure:"logMode" yaml:"logMode"`                         // 日志级别
	EnableFileLogWriter bool   `mapstructure:"enableFileLogWriter" yaml:"enableFileLogWriter"` // 是否启用日志文件
	LogFolder           string `mapstructure:"logFolder" yaml:"logFolder"`                     // 日志文件名称
}

// Redis Redis配置
type Redis struct {
	Host     string `mapstructure:"host" yaml:"host"`         // 主机地址
	Port     string `mapstructure:"port" yaml:"port"`         // 端口号
	Db       int    `mapstructure:"db" yaml:"db"`             // 库名
	Password string `mapstructure:"password" yaml:"password"` // 密码
}
