package entity

import "time"

/**
 * 地址断言表
 * @WAIT: 第一版先不写
 * @describe 仅在微服务模式下有效
 * @author eyesYeager
 * @date 2023/4/14 17:29
 */

type Predicate struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	ServerId   string    `json:"server_id" gorm:"not null;comment:服务id"`
	Path       string    `json:"path" gorm:"comment:地址断言，仅在微服务模式下有效"`
	Status     uint8     `json:"status" gorm:"not null;default:1;comment:1(正常),2(停用)"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:备注"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"default:null"`
}
