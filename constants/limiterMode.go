package constants

/**
 * 限流器模式
 * @auth eyesYeager
 * @date 2024/3/7 15:44
 */

var LimiterMode = limiterMode{
	Global: "global", // 全局限流
	Ip:     "ip",     // 基于ip限流
}

type limiterMode struct {
	Global string
	Ip     string
}
