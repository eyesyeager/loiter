package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/3/7 15:18
 */

type GetLimiterByPage struct {
	AppId       uint      `json:"appId"`
	AppName     string    `json:"appName"`
	Mode        string    `json:"mode"`
	LimiterCode string    `json:"limiterCode"`
	LimiterName string    `json:"limiterName"`
	Parameter   string    `json:"parameter"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
