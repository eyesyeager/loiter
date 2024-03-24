package receiver

/**
 * 添加黑白名单ip
 * @auth eyesYeager
 * @date 2024/1/24 19:11
 */

type AddNameListIp struct {
	AppId  uint     `json:"appId" validate:"required"`  // 应用id
	Genre  string   `json:"genre" validate:"required"`  // 名单类型，可选值见 constants.NameList
	IpList []string `json:"ipList" validate:"required"` // ip列表
}
