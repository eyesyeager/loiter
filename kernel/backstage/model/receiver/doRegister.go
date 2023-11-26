package receiver

/**
 * @author eyesYeager
 * @date 2023/11/26 16:02
 */

type DoRegister struct {
	Username string `json:"username" validate:"required"` // 用户名
	Email    string `json:"email" validate:"required"`    // 邮箱
	RoleId   uint   `json:"role_id" validate:"required"`  // 角色id
	Remarks  string `json:"remarks"`                      // 备注
}
