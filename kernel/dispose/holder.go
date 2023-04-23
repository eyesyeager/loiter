package dispose

/**
 * 代理配置容器
 * @author eyesYeager
 * @date 2023/4/23 15:00
 */

// --------------------------------------- AppHolder ------------------------------------------------ //

// AppHolder app容器
var AppHolder = map[string]AppHolderModel{}

// AppHolderModel APP容器模型
type AppHolderModel struct {
	Pattern  uint8
	Balancer uint8
	Servers  []ServerHolderModel
}

// ServerHolderModel 服务容器模型
type ServerHolderModel struct {
	Server string
	Weight uint
	Group  string
}

// --------------------------------------- XXHolder ------------------------------------------------ //

// 过滤器容器
// 初始化或更新过滤器容器
