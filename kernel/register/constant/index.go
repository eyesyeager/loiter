package constant

/**
 * @author eyesYeager
 * @date 2023/4/26 17:21
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
