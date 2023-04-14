package entity

import "time"

/**
 * 操作日志
 * @author eyesYeager
 * @date 2023/4/14 14:16
 */

type LogOperate struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Pattern    string    `json:"pattern" gorm:"not null;comment:操作类型"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:详细信息"`
	Ip         string    `json:"ip" gorm:"not null"`
	Browser    string    `json:"browser" gorm:"not null"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
