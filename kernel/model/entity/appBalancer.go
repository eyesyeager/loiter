package entity

import "gorm.io/gorm"

/**
 * 应用-负载均衡策略表
 * @auth eyesYeager
 * @date 2024/1/5 16:37
 */

type AppBalancer struct {
	gorm.Model
	AppId      uint `json:"appId" gorm:"not null;unique;comment:应用id"`
	BalancerId uint `json:"balancerId" gorm:"not null;comment:策略id"`
}
