package entity

import "gorm.io/gorm"

/**
 * 访问日志表
 * @auth eyesYeager
 * @date 2024/1/11 17:06
 */

type RequestLog struct {
	gorm.Model
	Host    string `json:"host" gorm:"not null;comment:主机地址"`
	Path    string `json:"path" gorm:"comment:请求接口"`
	ReqInfo string `json:"req_info" gorm:"type:TEXT;comment:请求信息"`
	Ip      string `json:"ip" gorm:"comment:IP地址"`
	Browser string `json:"browser" gorm:"浏览器"`
}
