package receiver

import (
	"loiter/model/structure"
)

/**
 * 获取通用日志
 * @auth eyesYeager
 * @date 2024/1/4 15:15
 */

type GetUniversalLog struct {
	structure.PageStruct        // 分页参数
	OperatorName         string `json:"operatorName"` // 操作人
	Title                string `json:"title"`        // 标题
	Content              string `json:"content"`      // 内容
	TimeBegin            string `json:"timeBegin"`    // 时间从
	TimeEnd              string `json:"timeEnd"`      // 时间至
}
