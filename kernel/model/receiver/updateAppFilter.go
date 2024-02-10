package receiver

/**
 * 更新应用过滤器
 * @auth eyesYeager
 * @date 2024/1/11 18:06
 */

type UpdateAppFilter struct {
	AppId           uint     `json:"appId" validate:"required"`           // 应用id
	FilterNameSlice []string `json:"filterNameSlice" validate:"required"` // 应用过滤器名切片
}
