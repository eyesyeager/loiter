package receiver

/**
 * 更新应用处理器
 * @auth eyesYeager
 * @date 2024/1/11 18:06
 */

type SaveAppProcessor struct {
	AppId     uint     `json:"appId" validate:"required"`     // 应用id
	Filter    []string `json:"filter" validate:"required"`    // 过滤器
	Aid       []string `json:"aid" validate:"required"`       // 后置处理器
	Exception []string `json:"exception" validate:"required"` // 异常处理器
	Final     []string `json:"final" validate:"required"`     // 最终处理器
}
