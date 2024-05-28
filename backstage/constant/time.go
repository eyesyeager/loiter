package constant

/**
 * 时间常量
 * @auth eyesYeager
 * @date 2024/2/17 13:53
 */

var Time = time{
	HourSeconds: 60 * 60,
	DaySeconds:  60 * 60 * 24,
}

type time struct {
	HourSeconds int
	DaySeconds  int
}
