package constant

/**
 * 通道表
 * @auth eyesYeager
 * @date 2024/1/9 19:22
 */

var Passageway = passageway{
	1,
	2,
}

type passageway struct {
	Filter   uint8
	Pipeline uint8
}
