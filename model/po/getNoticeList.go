package po

import "time"

/**
 * @auth eyesYeager
 * @date 2024/3/11 17:50
 */

type GetNoticeList struct {
	Id        uint
	AppName   string
	Host      string
	Title     string
	Genre     string
	Content   string
	Remarks   string
	CreatedAt time.Time
}
