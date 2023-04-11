package config

/**
 * 程序运行相关配置
 * @author eyesYeager
 * @date 2023/4/10 17:23
 */

type Program struct {
	Key string `mapstructure:"key" yaml:"key"` // 网关密钥
}
