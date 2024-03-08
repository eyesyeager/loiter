package receiver

import "loiter/model/structure"

/**
 * 分页获取应用信息
 * @auth eyesYeager
 * @date 2024/2/21 16:23
 */

type GetAppInfoByPage struct {
	structure.PageStruct        // 分页参数
	AppName              string `json:"appName"` // 应用名
	Status               string `json:"status"`  // 应用状态
}
