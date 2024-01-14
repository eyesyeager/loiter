package receiver

/**
 * 注册应用实例
 * @auth eyesYeager
 * @date 2024/1/5 14:05
 */

type AddServer struct {
	AppId   uint   `json:"appId" validate:"required"`   // 关联应用id
	Name    string `json:"name"`                        // 应用实例名
	Address string `json:"address" validate:"required"` // 地址
	Remarks string `json:"remarks"`                     // 备注
}
