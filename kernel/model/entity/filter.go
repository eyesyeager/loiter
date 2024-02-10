package entity

import "gorm.io/gorm"

/**
 * 过滤器配置表
 * @auth eyesYeager
 * @date 2024/1/9 19:19
 */

type Filter struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;comment:过滤器名"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
