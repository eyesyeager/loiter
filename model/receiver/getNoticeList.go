package receiver

import "loiter/model/structure"

/**
 * @auth eyesYeager
 * @date 2024/2/23 11:16
 */

type GetNoticeList struct {
	structure.PageStruct        // 分页参数
	AppName              string `json:"appName"`   // 应用名
	Title                string `json:"title"`     // 标题
	Genre                string `json:"genre"`     // 类型
	TimeBegin            string `json:"timeBegin"` // 时间从
	TimeEnd              string `json:"timeEnd"`   // 时间至
}
