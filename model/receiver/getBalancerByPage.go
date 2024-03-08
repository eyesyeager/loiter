package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/2/29 19:35
 */

type GetBalancerByPage struct {
	structure.PageStruct        // 分页参数
	AppName              string `json:"appName"`  // 应用名
	Balancer             string `json:"balancer"` // 负载均衡策略
}
