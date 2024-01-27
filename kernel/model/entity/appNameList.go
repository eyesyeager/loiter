package entity

import "gorm.io/gorm"

/**
 * 应用-黑白名单表
 * @auth eyesYeager
 * @date 2024/1/24 11:28
 */

type AppNameList struct {
	gorm.Model
	AppId uint   `json:"appId" gorm:"not null;comment:应用id"`
	Genre string `json:"genre" gorm:"not null;comment:名单类型，可选值为：black，white"`
}
