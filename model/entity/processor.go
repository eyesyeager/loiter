package entity

import "gorm.io/gorm"

/**
 * 处理器表
 * @auth eyesYeager
 * @date 2024/2/12 18:52
 */

type Processor struct {
	gorm.Model
	Code    string `json:"code" gorm:"not null;unique;comment:处理器编码"`
	Name    string `json:"name" gorm:"not null;unique;comment:处理器名"`
	Genre   string `json:"genre" gorm:"not null;comment:类型，见constant.Processor"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
