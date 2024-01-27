package entity

import "gorm.io/gorm"

/**
 * 应用表
 * @author eyesYeager
 * @date 2023/4/11 15:21
 */

type App struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;unique;comment:应用名称"`
	Host    string `json:"host" gorm:"not null;unique;comment:应用地址"`
	Status  uint8  `json:"status" gorm:"not null;default:2;comment:1(正常),2(停用)"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
