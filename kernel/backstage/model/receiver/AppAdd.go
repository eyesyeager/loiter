package receiver

/**
 * 添加app
 * @author eyesYeager
 * @date 2023/4/12 10:39
 */

type AppAdd struct {
	App     string `json:"app" validate:"required"`     // 应用地址
	Pattern uint8  `json:"pattern" validate:"required"` // 模式
	Remarks string `json:"remarks"`                     // 备注
}
