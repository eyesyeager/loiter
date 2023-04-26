package balancer

/**
 * @author eyesYeager
 * @date 2023/4/26 17:21
 */

// LoadPattern 负载均衡模式
var LoadPattern = map[string]uint8{
	"random": 1, // 随机
	"round":  2, // 轮询
	"weight": 3, // 加权轮询
}
