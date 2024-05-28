package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/3/7 14:32
 */

type GetLimiterByPage struct {
	structure.PageStruct        // 分页参数
	AppId                uint   `json:"appId"`   // 应用id
	Limiter              string `json:"limiter"` // 限流器类型
}
