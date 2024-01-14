package returnee

import (
	"loiter/kernel/model/structure"
)

/**
 * 获取登录日志
 * @auth eyesYeager
 * @date 2024/1/3 19:20
 */

type GetLoginLog struct {
	structure.PageStruct                    // 分页参数
	Total                int64              // 总数
	Data                 []GetLoginLogInner // 数据
}

type GetLoginLogInner struct {
	Username  string `json:"username"`  // 登录人
	Ip        string `json:"ip"`        // ip
	Browser   string `json:"browser"`   // 浏览器
	CreatedAt string `json:"createdAt"` // 创建时间
}
