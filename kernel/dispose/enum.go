package dispose

/**
 * 全局枚举变量
 * @author eyesYeager
 * @date 2023/4/11 17:52
 */

// AppStatus App状态
var AppStatus = map[string]uint8{
	"normal":     1, // 正常
	"deactivate": 2, // 停用
}

// AppPattern APP模式
var AppPattern = map[string]uint8{
	"singleton": 1, // 单体
	"micro":     2, // 微服务
}

// ServerStatus 服务状态
var ServerStatus = map[string]uint8{
	"normal":     1, // 正常
	"deactivate": 2, // 停用
}

// BalancerPattern 负载均衡模式
var BalancerPattern = map[string]uint8{
	"random": 1, // 随机
	"round":  2, // 轮询
	"weight": 3, // 带权重的轮询
}
