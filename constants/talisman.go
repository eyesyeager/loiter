package constants

/**
 * 特殊标识
 * @auth eyesYeager
 * @date 2024/2/23 14:35
 */

var Talisman = talisman{
	WithoutApp: "withoutApp", // 操作不针对应用
}

type talisman struct {
	WithoutApp string
}
