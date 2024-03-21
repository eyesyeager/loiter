package receiver

/**
 * @author eyesYeager
 * @date 2024/3/21 14:44
 */

type SaveStaticApp struct {
	AppId      uint   `json:"appId" validate:"required"`
	ErrorRoute string `json:"errorRoute"`
}
