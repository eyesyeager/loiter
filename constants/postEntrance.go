package constants

/**
 * post代理入口方式
 * @auth eyesYeager
 * @date 2024/2/16 13:25
 */

var PostEntrance = postEntrance{
	Error:  "err",
	Reject: "reject",
	Post:   "post",
}

type postEntrance struct {
	Error  string
	Reject string
	Post   string
}
