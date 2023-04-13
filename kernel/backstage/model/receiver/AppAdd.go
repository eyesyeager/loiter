package receiver

/**
 * @author eyesYeager
 * @date 2023/4/12 10:39
 */

type AppAdd struct {
	Host    string `json:"host" validate:"required"`    // 主机地址
	Pattern int8   `json:"pattern" validate:"required"` // 模式
	Status  int8   `json:"status" validate:"required"`  // 状态
}
