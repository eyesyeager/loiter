package entity

import "gorm.io/gorm"

/**
 * 应用-响应处理器
 * @auth eyesYeager
 * @date 2024/1/26 09:46
 */

type AppAid struct {
	gorm.Model
	AppId   uint   `json:"app_id" gorm:"not null;unique;comment:应用名"`
	AidName string `json:"aid_name" gorm:"not null;comment:用符号有序拼接的响应处理器名"`
}
