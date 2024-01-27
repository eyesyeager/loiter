package plugin

import (
	"loiter/plugin/aid"
	"loiter/plugin/loadbalancer"
	"loiter/plugin/passageway"
)

/**
 * 应用插件配置通道
 * @auth eyesYeager
 * @date 2024/1/22 11:49
 */

// Register 注册插件
func Register() {
	// 注册负载均衡插件
	loadbalancer.Register()
	// 注册通道插件
	passageway.Register()
	// 注册响应处理器插件
	aid.Register()
}
