package balancer

/**
 * 负载均衡策略
 * @author eyesYeager
 * @date 2023/4/23 15:10
 */

// StrategyBalancer 策略类型
//type StrategyBalancer func(servers []dispose.ServerHolderModel) string

// StrategyMap 负载均衡策略
//var StrategyMap = map[uint8]StrategyBalancer{
//	LoadPattern["random"]: RandomBalancer,        // 随机
//	LoadPattern["round"]:  RoundBalancer,         // 轮询
//	LoadPattern["weight"]: WeightedRoundBalancer, // 带权重的轮询
//}

// ----------------------------------------- 策略的具体实现 -----------------------------------------------------

// RandomBalancer 随机策略
//func RandomBalancer(servers []dispose.ServerHolderModel) string {
//	// 我并没有给它一个随机的种子，因为我认为这不重要，而且设置种子的操作会带来性能损失
//	index := rand.Intn(len(servers))
//	return servers[index].Server
//}

// RoundBalancer 轮询策略
//func RoundBalancer(servers []dispose.ServerHolderModel) string {
//	return ""
//}

// WeightedRoundBalancer 加权轮询策略
//func WeightedRoundBalancer(servers []dispose.ServerHolderModel) string {
//	return ""
//}
