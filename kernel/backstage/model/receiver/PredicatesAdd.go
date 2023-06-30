package receiver

/**
 * 添加Predicates
 * @author eyesYeager
 * @date 2023/4/26 23:00
 */

type PredicatesAdd struct {
	BasketId uint   `json:"basketId" validate:"required"` // 组id
	Path     string `json:"path" validate:"required"`     // 断言路径
	NeedPath bool   `json:"needPath"`                     // 代理时是否需要保留断言路径
	Remarks  string `json:"remarks"`                      // 备注
}
