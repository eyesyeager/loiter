package config

/**
 * 运行环境配置
 * @author eyesYeager
 * @date 2023/4/10 15:53
 */

type Profiles struct {
	Active string `mapstructure:"active" yaml:"active"` // 运行环境
}
