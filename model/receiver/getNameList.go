package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/5/14 10:37
 */

type GetNameList struct {
	structure.PageStruct        // 分页参数
	AppId                uint   `json:"appId"`     // 应用Id
	Genre                string `json:"genre"`     // 类型，可选值见 constants.NameList
	Ip                   string `json:"ip"`        // ip
	Remarks              string `json:"remarks"`   // 备注
	TimeBegin            string `json:"timeBegin"` // 时间从
	TimeEnd              string `json:"timeEnd"`   // 时间至
}
