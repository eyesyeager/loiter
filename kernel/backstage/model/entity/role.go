package entity

import (
	"gorm.io/gorm"
)

/**
 * 角色表
 * @author eyesYeager
 * @date 2023/9/26 14:44
 */

type Role struct {
	gorm.Model
	Name    string `json:"name" gorm:"comment:角色名"`
	Weight  uint   `json:"weight" gorm:"default:1;comment:角色权重"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
