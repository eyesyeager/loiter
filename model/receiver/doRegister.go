package receiver

/**
 * 开通新账号请求参数
 * @author eyesYeager
 * @date 2023/11/26 16:02
 */

type DoRegister struct {
	Username string `json:"username" validate:"required"` // 用户名
	Email    string `json:"email" validate:"required"`    // 邮箱
	Role     string `json:"role" validate:"required"`     // 角色名
	Remarks  string `json:"remarks"`                      // 备注
}
