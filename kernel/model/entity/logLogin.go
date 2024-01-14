package entity

import "gorm.io/gorm"

/**
 * 登录日志表
 * @author eyesYeager
 * @date 2023/9/27 12:31
 */

type LogLogin struct {
	gorm.Model
	Uid     uint   `json:"user_id" gorm:"not null;comment:用户id"`
	Ip      string `json:"ip" gorm:"not null;comment:ip地址"`
	Browser string `json:"browser" gorm:"comment:浏览器"`
}
