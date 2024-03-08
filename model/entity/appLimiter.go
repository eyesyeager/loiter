package entity

import "gorm.io/gorm"

/**
 * 应用-限流策略表
 * @auth eyesYeager
 * @date 2024/1/12 16:06
 */

type AppLimiter struct {
	gorm.Model
	AppId     uint   `json:"appId" gorm:"not null;unique;comment:应用id"`
	Limiter   string `json:"limiter" gorm:"not null;comment:限流策略"`
	Parameter string `json:"parameter" gorm:"comment:json格式的参数设置"`
}
