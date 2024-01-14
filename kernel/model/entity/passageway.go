package entity

import "gorm.io/gorm"

/**
 * 通道配置表
 * @auth eyesYeager
 * @date 2024/1/9 19:19
 */

type Passageway struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;comment:通道名"`
	Genre   string `json:"genre" gorm:"not null;comment:通道类型，见constant.Passageway"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
