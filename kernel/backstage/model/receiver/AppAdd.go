package receiver

/**
 * @author eyesYeager
 * @date 2023/4/12 10:39
 */

type AppAdd struct {
	App     string `json:"app" validate:"required"`     // 应用地址
	Pattern int8   `json:"pattern" validate:"required"` // 模式
	Status  int8   `json:"status" validate:"required"`  // 状态
}
