package entity

import "gorm.io/gorm"

/**
 * 应用-处理器表
 * @auth eyesYeager
 * @date 2024/2/12 20:23
 */

type AppProcessor struct {
	gorm.Model
	AppId uint   `json:"app_id" gorm:"not null;comment:应用id"`
	Genre string `json:"genre" gorm:"not null;comment:类型，见constant.Processor"`
	Codes string `json:"codes" gorm:"not null;comment:用符号有序拼接的处理器编码"`
}
