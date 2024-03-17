package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/3/7 14:32
 */

type GetLimiterByPage struct {
	structure.PageStruct        // 分页参数
	AppName              string `json:"appName"` // 应用名
	Limiter              string `json:"limiter"` // 限流器类型
}
