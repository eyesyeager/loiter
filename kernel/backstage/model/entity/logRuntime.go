package entity

import "time"

/**
 * 运行时日志
 * @author eyesYeager
 * @date 2023/4/14 14:30
 */

type LogRuntime struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:操作类型"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:详细信息"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
}
