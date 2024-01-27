package receiver

/**
 * 添加黑白名单ip
 * @auth eyesYeager
 * @date 2024/1/24 19:11
 */

type AddNameListIp struct {
	AppId      uint   `json:"appId" validate:"required"`      // 应用id
	Genre      string `json:"genre" validate:"required"`      // 名单类型，可选值为：white、black
	IpSliceStr string `json:"ipSliceStr" validate:"required"` // ip拼接而成的字符串
}
