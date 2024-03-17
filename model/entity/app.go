package entity

import "gorm.io/gorm"

/**
 * 应用表
 * @author eyesYeager
 * @date 2023/4/11 15:21
 */

type App struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;comment:应用名称"`
	Host    string `json:"host" gorm:"not null;unique;comment:应用地址"`
	Genre   string `json:"genre" gorm:"not null;comment:应用类型，见constants.AppGenre"`
	OwnerId uint   `json:"owner_id" gorm:"not null;comment:应用负责人id"`
	Status  uint8  `json:"status" gorm:"default:2;comment:1(正常),2(停用)"`
	Remarks string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
