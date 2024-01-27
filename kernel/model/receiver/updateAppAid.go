package receiver

/**
 * 更新应用响应处理器
 * @auth eyesYeager
 * @date 2024/1/26 16:15
 */

type UpdateAppAid struct {
	AppId        uint     `json:"appId" validate:"required"`        // 应用id
	AidNameSlice []string `json:"aidNameSlice" validate:"required"` // 应用响应处理器名切片
}
