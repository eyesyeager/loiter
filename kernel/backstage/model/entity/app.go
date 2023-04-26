package entity

import "time"

/**
 * 应用表
 * @author eyesYeager
 * @date 2023/4/11 15:21
 */

type App struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	App        string    `json:"app" gorm:"not null;unique;comment:应用地址"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:1(单体),2(微服务)"`
	Status     uint8     `json:"status" gorm:"not null;default:2;comment:1(正常),2(停用)"`
	Remarks    string    `json:"remarks" gorm:"type:TEXT;comment:备注"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"default:null"`
}
