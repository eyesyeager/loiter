package receiver

/**
 * 注册应用
 * @auth eyesYeager
 * @date 2024/1/4 18:03
 */

type SaveApp struct {
	AppId      uint         `json:"appId"`                          // 应用id
	AppName    string       `json:"appName" validate:"required"`    // 应用名
	AppGenre   string       `json:"appGenre" validate:"required"`   // 应用类型
	Host       string       `json:"host" validate:"required"`       // 地址
	OwnerId    uint         `json:"ownerId" validate:"required"`    // 责任人id
	ServerList []SaveServer `json:"serverList" validate:"required"` // 实例列表
	Remarks    string       `json:"remarks"`                        // 备注
}

type SaveServer struct {
	Address string `json:"address"` // 服务实例地址
	Weight  uint   `json:"weight"`  // 权重
}
