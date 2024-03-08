package returnee

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/2/29 19:43
 */

type GetBalancerByPage struct {
	structure.PageStruct                          // 分页参数
	Total                int64                    `json:"total"` // 总数
	Data                 []GetBalancerByPageInner `json:"data"`  // 数据
}

type GetBalancerByPageInner struct {
	Id           uint   `json:"id"`           // 应用负载均衡id
	AppName      string `json:"appName"`      // 应用名
	BalancerName string `json:"balancerName"` // 负载均衡策略名
	BalancerCode string `json:"balancerCode"` // 负载均衡策略代码
	Operator     string `json:"operator"`     // 操作人
	UpdatedAt    string `json:"updatedAt"`    // 操作时间
}
