package entity

import "gorm.io/gorm"

/**
 * 访问日志表
 * @auth eyesYeager
 * @date 2024/1/11 17:06
 */

type RequestLog struct {
	gorm.Model
	RequestId string `json:"request_id" gorm:"not null;comment:请求唯一标识"`
	Host      string `json:"host" gorm:"not null;comment:主机地址"`
	Path      string `json:"path" gorm:"comment:请求接口"`
	ReqInfo   string `json:"req_info" gorm:"type:TEXT;comment:请求信息"`
	Ip        string `json:"ip" gorm:"comment:IP地址"`
	Browser   string `json:"browser" gorm:"comment:浏览器"`
	StartTime string `json:"start_time" gorm:"comment:请求开始时间戳(ms)"`
	EndTime   string `json:"end_time" gorm:"comment:请求结束时间戳(ms)"`
	RunTime   uint64 `json:"run_time" gorm:"comment:请求耗时(ms)"`
	Entrance  string `json:"entrance" gorm:"comment:日志入口"`
	ErrInfo   string `json:"err_info" gorm:"comment:错误信息"`
}
