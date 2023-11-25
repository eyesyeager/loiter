package constant

/**
 * @author eyesYeager
 * @date 2023/9/28 9:14
 */

var Status = status{
	1,
	2,
	3,
}

type status struct {
	Normal  uint
	Invalid uint
	Delete  uint
}
