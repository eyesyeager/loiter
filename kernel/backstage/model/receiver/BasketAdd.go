package receiver

/**
 * 添加basket
 * @author eyesYeager
 * @date 2023/4/26 16:07
 */

type BasketAdd struct {
	AppId    uint   `json:"appId" validate:"required"` // 应用id
	Name     string `json:"name" validate:"required"`  // 组名
	Balancer uint8  `json:"balancer"`                  // 负载均衡策略
	Remarks  string `json:"remarks"`                   // 备注
}
