package receiver

/**
 * 注册应用
 * @auth eyesYeager
 * @date 2024/1/4 18:03
 */

type AddApp struct {
	Name    string `json:"name" validate:"required"` // 应用名
	Host    string `json:"host" validate:"required"` // 地址
	Remarks string `json:"remarks"`                  // 备注
}
