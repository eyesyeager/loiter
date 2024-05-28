package receiver

/**
 * @auth eyesYeager
 * @date 2024/3/8 10:51
 */

type DeleteAppLimiter struct {
	AppId uint `json:"appId" validate:"required"`
}
