package returnee

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/5/14 10:47
 */

type GetNameList struct {
	structure.PageStruct                    // 分页参数
	Total                int64              `json:"total"` // 总数
	Data                 []GetNameListInner `json:"data"`  // 数据
}

type GetNameListInner struct {
	Id        uint   `json:"id"`        // 名单id
	AppName   string `json:"appName"`   // 应用名
	AppId     uint   `json:"appId"`     // 应用id
	Ip        string `json:"ip"`        // ip
	Remarks   string `json:"remarks"`   // 备注
	CreatedAt string `json:"createdAt"` // 创建时间
}
