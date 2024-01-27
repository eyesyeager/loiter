package entity

import "gorm.io/gorm"

/**
 * 应用-通道表
 * @auth eyesYeager
 * @date 2024/1/9 19:28
 */

type AppPassageway struct {
	gorm.Model
	AppId          uint   `json:"app_id" gorm:"not null;unique;comment:应用id"`
	PassagewayName string `json:"passageway_name" gorm:"not null;comment:用符号有序拼接的通道名"`
}
