package entity

import "time"

/**
 * 服务表
 * @author eyesYeager
 * @date 2023/4/13 15:40
 */

type Server struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	AppId      uint      `json:"app_id" gorm:"not null;comment:应用id"`
	Server     string    `json:"server" gorm:"not null;comment:服务地址"`
	Status     uint8     `json:"status" gorm:"not null;default:1;comment:1(正常),2(停用)"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:备注"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"default:null"`
}
