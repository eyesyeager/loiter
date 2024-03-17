package returnee

import (
	"loiter/model/structure"
)

/**
 * @auth eyesYeager
 * @date 2024/2/23 11:25
 */

type GetNoticeList struct {
	structure.PageStruct                      // 分页参数
	Total                int64                `json:"total"` // 总数
	Data                 []GetNoticeListInner `json:"data"`  // 数据
}

type GetNoticeListInner struct {
	Id        uint   `json:"id"`        // 通知id
	AppName   string `json:"appName"`   // 应用名
	Title     string `json:"title"`     // 标题
	Genre     string `json:"genre"`     // 类型
	Content   string `json:"content"`   // 内容
	Remarks   string `json:"remarks"`   // 备注
	CreatedAt string `json:"createdAt"` // 创建时间
}
