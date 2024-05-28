package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/5/14 11:21
 */

type GetNameList struct {
	Id        uint      `json:"id"`        // 名单id
	AppName   string    `json:"appName"`   // 应用名
	AppId     uint      `json:"appId"`     // 应用id
	Ip        string    `json:"ip"`        // ip
	Remarks   string    `json:"remarks"`   // 备注
	CreatedAt time.Time `json:"createdAt"` // 创建时间
}
