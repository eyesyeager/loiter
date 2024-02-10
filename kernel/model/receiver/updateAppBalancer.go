package receiver

/**
 * 更新负载均衡策略请求参数
 * @auth eyesYeager
 * @date 2024/1/5 16:46
 */

type UpdateAppBalancer struct {
	AppId      uint `json:"appId" validate:"required"`      // 应用id
	BalancerId uint `json:"balancerId" validate:"required"` // 负载均衡策略id
}
