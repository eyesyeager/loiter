package entity

import (
	"gorm.io/gorm"
)

/**
 * 登录日志表
 * @author eyesYeager
 * @date 2023/9/27 12:31
 */

type LogLogin struct {
	gorm.Model
	UserId  uint   `json:"user_id" gorm:"not null"`
	Ip      string `json:"ip" gorm:"not null"`
	Os      string `json:"os"`
	Browser string `json:"browser"`
}
