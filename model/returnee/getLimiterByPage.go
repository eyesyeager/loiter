package returnee

import (
	"loiter/model/structure"
)

/**
 * @auth eyesYeager
 * @date 2024/3/7 14:42
 */

type GetLimiterByPage struct {
	structure.PageStruct                         // 分页参数
	Total                int64                   `json:"total"` // 总数
	Data                 []GetLimiterByPageInner `json:"data"`  // 数据
}

type GetLimiterByPageInner struct {
	AppId       uint   `json:"appId"`
	AppName     string `json:"appName"`
	Mode        string `json:"mode"`
	LimiterCode string `json:"limiterCode"`
	LimiterName string `json:"limiterName"`
	Parameter   string `json:"parameter"`
	UpdatedAt   string `json:"updatedAt"`
}
