package returnee

import (
	"loiter/model/structure"
)

/**
 * 获取通用日志
 * @auth eyesYeager
 * @date 2024/1/4 15:15
 */

type GetUniversalLog struct {
	structure.PageStruct                        // 分页参数
	Total                int64                  `json:"total"` // 总数
	Data                 []GetUniversalLogInner `json:"data"`  // 数据
}

type GetUniversalLogInner struct {
	Operator  string `json:"operator"`  // 操作人
	Title     string `json:"title"`     // 日志标题
	Content   string `json:"content"`   // 日志内容
	Ip        string `json:"ip"`        // ip
	Browser   string `json:"browser"`   // 浏览器
	CreatedAt string `json:"createdAt"` // 创建时间
}
