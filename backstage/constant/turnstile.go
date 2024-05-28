package constant

/**
 * 开关
 * @auth eyesYeager
 * @date 2024/1/24 17:46
 */

var Turnstile = turnstile{
	Open:  1,
	Close: 2,
}

type turnstile struct {
	Open  int8
	Close int8
}

// CheckTurnstile 检查开关值是否合法
func CheckTurnstile(value int8) bool {
	return value == Turnstile.Open || value == Turnstile.Close
}

// GetTurnstileName 根据值获取开关名字
func GetTurnstileName(value int8) string {
	if value == Turnstile.Open {
		return "open"
	}
	if value == Turnstile.Close {
		return "close"
	}
	return ""
}
