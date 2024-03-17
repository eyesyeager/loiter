package receiver

/**
 * 更新应用限流器
 * @auth eyesYeager
 * @date 2024/1/23 10:30
 */

type SaveAppLimiter struct {
	AppId     uint   `json:"appId" validate:"required"`   // 应用id
	Limiter   string `json:"limiter" validate:"required"` // 限流器
	Mode      string `json:"mode" validate:"required"`    // 限流模式
	Parameter string `json:"parameter"`                   // 参数
}
