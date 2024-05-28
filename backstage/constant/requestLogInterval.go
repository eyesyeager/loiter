package constant

/**
 * 请求日志间隔
 * @auth eyesYeager
 * @date 2024/2/15 17:40
 */

var RequestLogInterval = requestLogInterval{
	Today:     "today",
	Yesterday: "yesterday",
	Week:      "week",
	Month:     "month",
}

type requestLogInterval struct {
	Today     string
	Yesterday string
	Week      string
	Month     string
}
