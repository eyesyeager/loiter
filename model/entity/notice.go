package entity

import "gorm.io/gorm"

/**
 * 通知表
 * @auth eyesYeager
 * @date 2024/2/20 15:05
 */

type Notice struct {
	gorm.Model
	AppName string `json:"app_name" gorm:"comment:应用名"`
	Title   string `json:"title" gorm:"comment:消息标题"`
	Content string `json:"content" gorm:"type:TEXT;comment:消息详情"`
	Genre   string `json:"genre" gorm:"not null;comment:类型，见constants.Notice"`
	Secret  bool   `json:"secret" gorm:"default:false;comment:是否保密"`
	Remarks string `json:"remarks" gorm:"comment:备注"`
}
