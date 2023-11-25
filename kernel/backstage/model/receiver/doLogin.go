package receiver

/**
 * @author eyesYeager
 * @date 2023/9/26 15:24
 */

type DoLogin struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}
