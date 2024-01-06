package receiver

/**
 * 更新负载均衡策略请求参数
 * @auth eyesYeager
 * @date 2024/1/5 16:46
 */

type UpdateAppBalance struct {
	AppID     uint `json:"appId" validate:"required"`     // 应用id
	BalanceId uint `json:"balanceId" validate:"required"` // 负载均衡策略id
}
