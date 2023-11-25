package entity

import "gorm.io/gorm"

/**
 * 用户表
 * @author eyesYeager
 * @date 2023/9/26 14:44
 */

type User struct {
	gorm.Model
	Username string `json:"name" gorm:"not null;unique;comment:用户名"`
	Password string `json:"password" gorm:"not null;comment:密码"`
	RoleId   uint   `json:"role" gorm:"not null;comment:角色id"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar" gorm:"comment:头像链接"`
	Status   uint8  `json:"status" gorm:"not null;default:1;comment:1(正常),2(停用),3(删除)"`
	Remarks  string `json:"remarks" gorm:"type:TEXT;comment:备注"`
}
