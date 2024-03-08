package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/2/22 09:34
 */

type GetAppInfoByPage struct {
	AppId     uint      `json:"appId"`
	AppName   string    `json:"appName"`
	Host      string    `json:"host"`
	Status    uint8     `json:"status"`
	Remarks   string    `json:"remarks"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	Balancer  string    `json:"balancer"`
	Plugins   int       `json:"plugins"`
}
