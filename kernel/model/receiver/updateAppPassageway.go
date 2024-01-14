package receiver

/**
 * 更新应用通道
 * @auth eyesYeager
 * @date 2024/1/11 18:06
 */

type UpdateAppPassageway struct {
	AppID               uint     `json:"appId" validate:"required"`               // 应用id
	PassagewayNameSlice []string `json:"passagewayNameSlice" validate:"required"` // 负载均衡策略id
}
