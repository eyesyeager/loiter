package receiver

/**
 * @author eyesYeager
 * @date 2023/4/13 15:50
 */

type AppServerAdd struct {
	AppId   uint   `json:"appId" validate:"required"`  // 应用id
	Server  string `json:"server" validate:"required"` // 服务地址
	Weight  uint   `json:"weight"`                     // 权重
	Group   string `json:"group"`                      // 服务所属组
	Status  int8   `json:"status"`                     // 状态
	Remarks string `json:"remarks"`                    // 备注
}
