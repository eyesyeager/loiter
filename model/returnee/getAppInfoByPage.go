package returnee

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/2/21 17:06
 */

type GetAppInfoByPage struct {
	structure.PageStruct                         // 分页参数
	Total                int64                   `json:"total"` // 总数
	Data                 []GetAppInfoByPageInner `json:"data"`  // 数据
}

type GetAppInfoByPageInner struct {
	AppId          uint   `json:"appId"`
	AppName        string `json:"appName"`
	Host           string `json:"host"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
	Owner          string `json:"owner"`
	CreatedAt      string `json:"createdAt"`
	ServerNum      int    `json:"serverNum"`
	ValidServerNum int    `json:"validServerNum"`
	Balancer       string `json:"balancer"`
	Plugins        int    `json:"plugins"`
}
