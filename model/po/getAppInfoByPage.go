package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/2/22 09:34
 */

type GetAppInfoByPage struct {
	AppId     uint      `json:"appId"`
	AppName   string    `json:"appName"`
	AppGenre  string    `json:"appGenre"`
	Host      string    `json:"host"`
	Status    uint8     `json:"status"`
	Remarks   string    `json:"remarks"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	Plugins   int       `json:"plugins"`
}
