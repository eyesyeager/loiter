package config

/**
 * 日志打印配置
 * @author eyesYeager
 * @date 2023/4/10 17:23
 */

type Log struct {
	Level           string `mapstructure:"level" yaml:"level"`                         // 日志等级
	RootDir         string `mapstructure:"root_dir" yaml:"root_dir"`                   // 日志根目录
	LogFolderApp    string `mapstructure:"log_folder_app" yaml:"log_folder_app"`       // 应用日志目录
	LogFolderServer string `mapstructure:"log_folder_server" yaml:"log_folder_server"` // 服务日志目录
	Format          string `mapstructure:"format" yaml:"format"`                       // 写入格式
	ShowLine        bool   `mapstructure:"show_line" yaml:"show_line"`                 // 是否显示调用行
	MaxBackups      int    `mapstructure:"max_backups" yaml:"max_backups"`             // 旧文件的最大个数
	MaxSize         int    `mapstructure:"max_size" yaml:"max_size"`                   // 日志文件最大大小（MB）
	MaxAge          int    `mapstructure:"max_age" yaml:"max_age"`                     // 旧文件的最大保留天数
	Compress        bool   `mapstructure:"compress" yaml:"compress"`                   // 是否压缩
}
