package entity

import "gorm.io/gorm"

/**
 * 应用实例
 * @auth eyesYeager
 * @date 2024/1/4 17:53
 */

type Server struct {
	gorm.Model
	AppId   uint   `json:"appId" gorm:"not null;comment:所属应用id"`
	Name    string `json:"name" gorm:"not null;comment:应用实例名称"`
	Address string `json:"address" gorm:"not null;comment:应用实例地址"`
	Weight  uint   `json:"weight" gorm:"not null;default:1;comment:权重"`
	Status  uint8  `json:"status" gorm:"not null;default:2;comment:1(正常),2(停用)"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
