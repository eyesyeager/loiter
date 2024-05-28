package entity

import "gorm.io/gorm"

/**
 * 黑白名单表
 * @auth eyesYeager
 * @date 2024/1/23 19:54
 */

type NameList struct {
	gorm.Model
	AppId   uint   `json:"appId" gorm:"not null;comment:应用id"`
	Genre   string `json:"genre" gorm:"not null;comment:名单类型，可选值为：black，white"`
	Ip      string `json:"ip" gorm:"not null;comment:ip"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
