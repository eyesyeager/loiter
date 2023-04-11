package config

/**
 * 应用程序基本配置
 * @author eyesYeager
 * @date 2023/4/10 17:14
 */

type App struct {
	Name          string `mapstructure:"name" yaml:"name"`                     // 应用名称
	Port          string `mapstructure:"port" yaml:"port"`                     // 网关服务端口号
	BackstagePort string `mapstructure:"backstage_port" yaml:"backstage_port"` // 后台管理系统服务端口号
}
