package entity

import "time"

/**
 * 服务组表
 * @author eyesYeager
 * @date 2023/4/25 11:30
 */

type Basket struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	AppId      uint      `json:"app_id" gorm:"not null;comment:应用id"`
	Name       string    `json:"name" gorm:"not null;comment:组名"`
	Balancer   uint8     `json:"balancer" gorm:"not null;default:1;comment:负载均衡算法"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"default:null"`
}
