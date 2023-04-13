package entity

import "time"

/**
 * 应用表
 * @author eyesYeager
 * @date 2023/4/11 15:21
 */

type App struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Host       string    `json:"host" gorm:"not null;comment:主机地址"`
	Pattern    uint8     `json:"pattern" gorm:"not null;comment:1(单机),2(微服务)"`
	Status     uint8     `json:"status" gorm:"not null;default:1;comment:1(正常),2(停用)"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	DeleteTime time.Time `json:"delete_time" gorm:"default:null;"`
}
