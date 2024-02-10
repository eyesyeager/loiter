package entity

import "gorm.io/gorm"

/**
 * 应用-过滤器表
 * @auth eyesYeager
 * @date 2024/1/9 19:28
 */

type AppFilter struct {
	gorm.Model
	AppId      uint   `json:"app_id" gorm:"not null;unique;comment:应用id"`
	FilterName string `json:"filter_name" gorm:"not null;comment:用符号有序拼接的过滤器名"`
}
