package constant

/**
 * 角色权值
 * @author eyesYeager
 * @date 2023/9/28 8:42
 */

var Role = role{
	100,
	50,
	1,
	0,
}

type role struct {
	SuperAdmin uint
	Admin      uint
	User       uint
	Visitor    uint
}
