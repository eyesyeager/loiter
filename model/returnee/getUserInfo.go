package returnee

/**
 * 获取用户信息
 * @auth eyesYeager
 * @date 2024/2/18 15:06
 */

type GetUserInfo struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	Weight   uint   `json:"weight"`
}
