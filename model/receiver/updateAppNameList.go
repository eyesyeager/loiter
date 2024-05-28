package receiver

/**
 * 更新应用黑白名单开启配置
 * @auth eyesYeager
 * @date 2024/1/24 17:11
 */

type UpdateAppNameList struct {
	AppId uint `json:"appId" validate:"required"` // 应用id
	Black bool `json:"black"`                     // 是否开启黑名单
	White bool `json:"white"`                     // 是否开启白名单
}
