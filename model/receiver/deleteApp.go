package receiver

/**
 * @auth eyesYeager
 * @date 2024/2/26 15:25
 */

type DeleteApp struct {
	AppId uint `json:"appId" validate:"required"`
}
