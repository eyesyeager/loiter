package receiver

import "loiter/kernel/backstage/model/structure"

/**
 * 获取通用日志
 * @auth eyesYeager
 * @date 2024/1/4 15:15
 */

type GetUniversalLog struct {
	structure.PageStruct        // 分页参数
	OperatorName         string `json:"operatorName"`   // 操作人
	Title                string `json:"title"`          // 标题
	Content              string `json:"content"`        // 内容
	LoginTimeBegin       string `json:"loginTimeBegin"` // 登录时间从
	LoginTimeEnd         string `json:"loginTimeEnd"`   // 登录时间至
}
