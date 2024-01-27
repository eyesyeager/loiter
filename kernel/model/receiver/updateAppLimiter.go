package receiver

/**
 * 更新应用限流器
 * @auth eyesYeager
 * @date 2024/1/23 10:30
 */

type UpdateAppLimiter struct {
	AppId       uint   `json:"appId" validate:"required"`       // 应用id
	LimiterName string `json:"limiterName" validate:"required"` // 限流器名称
	Parameter   string `json:"parameter"`                       // 参数
}
