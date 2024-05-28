package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/3/1 10:32
 */

type GetBalancerByPage struct {
	Id           uint      `json:"id"`           // 应用负载均衡id
	AppName      string    `json:"appName"`      // 应用名
	BalancerName string    `json:"balancerName"` // 负载均衡策略名
	BalancerCode string    `json:"balancerCode"` // 负载均衡策略代码
	Operator     string    `json:"operator"`     // 操作人
	UpdatedAt    time.Time `json:"updatedAt"`    // 操作时间
}
