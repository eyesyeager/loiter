package plugin

import (
	"loiter/plugin/aid"
	"loiter/plugin/balancer"
	"loiter/plugin/exception"
	"loiter/plugin/filter"
	"loiter/plugin/final"
)

/**
 * 应用插件配置
 * @auth eyesYeager
 * @date 2024/1/22 11:49
 */

// Register 注册插件
func Register() {
	// 注册过滤器插件
	filter.Register()
	// 注册负载均衡插件
	balancer.Register()
	// 注册响应处理器插件
	aid.Register()
	// 注册异常处理器插件
	exception.Register()
	// 注册最终处理器插件
	final.Register()
}
