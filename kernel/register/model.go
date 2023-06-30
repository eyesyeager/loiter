package register

/**
 * @author eyesYeager
 * @date 2023/4/27 11:28
 */

type AppModel struct {
}

// ServerModel 服务模型
type ServerModel struct {
	Server string // 服务地址
	Weight uint   // 服务权重
	Group  string // 服务集群
}
