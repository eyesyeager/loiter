package entity

import "gorm.io/gorm"

/**
 * 限流器表
 * @auth eyesYeager
 * @date 2024/1/12 16:01
 */

type Limiter struct {
	gorm.Model
	Code      string `json:"code" gorm:"not null;unique;comment:限流器编码"`
	Name      string `json:"name" gorm:"not null;comment:限流器名称"`
	Parameter string `json:"parameter" gorm:"comment:参数示例"`
	Remarks   string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
