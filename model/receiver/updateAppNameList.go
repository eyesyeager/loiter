package receiver

/**
 * 更新应用黑白名单开启配置
 * @auth eyesYeager
 * @date 2024/1/24 17:11
 */

type UpdateAppNameList struct {
	AppId     uint   `json:"appId" validate:"required"`     // 应用id
	Genre     string `json:"genre" validate:"required"`     // 名单类型，可选值为：white、black
	Turnstile int8   `json:"turnstile" validate:"required"` // 开关(取值见constant.Turnstile)
}
