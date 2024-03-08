package receiver

import (
	"loiter/model/structure"
)

/**
 * 查询登录日志请求参数
 * @auth eyesYeager
 * @date 2024/1/3 19:20
 */

type GetLoginLog struct {
	structure.PageStruct        // 分页参数
	Username             string `json:"username"`       // 登录人
	LoginTimeBegin       string `json:"loginTimeBegin"` // 登录时间从
	LoginTimeEnd         string `json:"loginTimeEnd"`   // 登录时间至
}
