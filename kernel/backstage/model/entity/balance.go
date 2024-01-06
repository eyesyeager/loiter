package entity

import "gorm.io/gorm"

/**
 * 负载均衡策略表
 * @auth eyesYeager
 * @date 2024/1/5 16:36
 */

type Balance struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;comment:策略名"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
