package entity

import "gorm.io/gorm"

/**
 * 通用日志表
 * @author eyesYeager
 * @date 2023/9/27 12:36
 */

type LogUniversal struct {
	gorm.Model
	Operator string `json:"operator" gorm:"not null;comment:操作人"`
	Ip       string `json:"ip" gorm:"not null;comment:ip地址"`
	Browser  string `json:"browser" gorm:"comment:浏览器"`
	Title    string `json:"title" gorm:"not null;comment:日志标题"`
	Content  string `json:"content" gorm:"not null;comment:日志内容"`
}
