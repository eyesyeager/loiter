package constants

/**
 * 消息类型
 * @auth eyesYeager
 * @date 2024/2/22 17:33
 */

var Notice = notice{
	Site:  "site",
	Email: "email",
}

type notice struct {
	Site  string
	Email string
}
