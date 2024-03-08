package po

/**
 * 获取用户登录角色
 * @author eyesYeager
 * @date 2023/9/28 9:12
 */

type LoginUserRole struct {
	Uid      uint
	Password string
	Role     string
	Status   uint8
}
