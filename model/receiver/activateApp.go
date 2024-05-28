package receiver

/**
 * @auth eyesYeager
 * @date 2024/2/29 14:26
 */

type ActivateApp struct {
	AppId uint `json:"appId" validate:"required"` // 应用id
}
