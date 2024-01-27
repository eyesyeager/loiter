package entity

import "gorm.io/gorm"

/**
 * 响应处理器
 * @auth eyesYeager
 * @date 2024/1/26 09:46
 */

type Aid struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;comment:响应处理器名"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
