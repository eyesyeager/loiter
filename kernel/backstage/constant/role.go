package constant

import "loiter/config"

/**
 * 角色权值
 * @author eyesYeager
 * @date 2023/9/28 8:42
 */

var Role = role{
	config.RoleConfig.SuperAdmin.Name,
	config.RoleConfig.Admin.Name,
	config.RoleConfig.User.Name,
	"visitor",
}

type role struct {
	SuperAdmin string
	Admin      string
	User       string
	Visitor    string
}
