package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/3/1 15:28
 */

type GetProcessorByPage struct {
	structure.PageStruct        // 分页参数
	AppName              string `json:"appName"` // 应用名
}
