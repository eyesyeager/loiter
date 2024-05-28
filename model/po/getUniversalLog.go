package po

import (
	"time"
)

/**
 * 获取通用日志
 * @auth eyesYeager
 * @date 2024/1/4 15:15
 */

type GetUniversalLogInner struct {
	Operator  string    `json:"operator"`  // 操作人
	Title     string    `json:"title"`     // 日志标题
	Content   string    `json:"content"`   // 日志内容
	Ip        string    `json:"ip"`        // ip
	Browser   string    `json:"browser"`   // 浏览器
	CreatedAt time.Time `json:"createdAt"` // 创建时间
}
