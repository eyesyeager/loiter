package receiver

/**
 * 获取请求日志详细信息
 * @auth eyesYeager
 * @date 2024/2/14 17:41
 */

type GetDetailedRequestLog struct {
	AppId        uint   `json:"appId"`                            // 应用名称
	TimeInterval string `json:"timeInterval" validate:"required"` // 间隔
}
