package receiver

import "loiter/model/structure"

/**
 * 分页获取应用信息
 * @auth eyesYeager
 * @date 2024/2/21 16:23
 */

type GetAppInfoByPage struct {
	structure.PageStruct        // 分页参数
	AppId                uint   `json:"appId"`    // 应用名
	AppGenre             string `json:"appGenre"` // 应用类型
	Status               uint8  `json:"status"`   // 应用状态
}
