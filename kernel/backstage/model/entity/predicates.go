package entity

import "time"

/**
 * 断言表
 * @author eyesYeager
 * @date 2023/4/26 15:06
 */

type Predicates struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	BasketId   uint      `json:"basket_id" gorm:"not null;comment:组id"`
	Path       string    `json:"path" gorm:"not null;comment:断言路径"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"default:null"`
}
