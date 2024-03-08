package returnee

/**
 * @auth eyesYeager
 * @date 2024/2/29 17:12
 */

type GetAppInfoById struct {
	AppId      uint        `json:"appId"`
	AppName    string      `json:"appName" validate:"required"`    // 应用名
	Host       string      `json:"host" validate:"required"`       // 地址
	OwnerId    uint        `json:"ownerId" validate:"required"`    // 责任人id
	ServerList []AppServer `json:"serverList" validate:"required"` // 实例列表
	Remarks    string      `json:"remarks"`                        // 备注
}

type AppServer struct {
	Address string `json:"address"` // 服务实例地址
	Weight  uint   `json:"weight"`  // 权重
}
