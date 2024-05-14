package receiver

/**
 * 删除黑白名单
 * @auth eyesYeager
 * @date 2024/1/24 19:55
 */

type DeleteNameListIp struct {
	Id    uint   `json:"id" validate:"required"`    // ip id
	Ip    string `json:"ip" validate:"required"`    // ip
	AppId uint   `json:"appId" validate:"required"` // 应用id
	Genre string `json:"genre" validate:"required"` // 名单类型，可选值为：white、black
}
