package po

import "time"

/**
 * 查询登录日志
 * @auth eyesYeager
 * @date 2024/1/3 19:20
 */

type GetLoginLogInner struct {
	Username  string    `json:"username"`  // 登录人
	Ip        string    `json:"ip"`        // ip
	Browser   string    `json:"browser"`   // 浏览器
	CreatedAt time.Time `json:"createdAt"` // 创建时间
}
