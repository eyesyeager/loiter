package constant

/**
 * 角色权值
 * @author eyesYeager
 * @date 2023/9/28 8:42
 */

var Role = role{
	"super_admin",
	"admin",
	"user",
	"visitor",
}

type role struct {
	SuperAdmin string
	Admin      string
	User       string
	Visitor    string
}
