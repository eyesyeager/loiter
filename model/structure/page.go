package structure

/**
 * 分页结构体
 * @auth eyesYeager
 * @date 2024/1/4 09:31
 */

type PageStruct struct {
	PageNo   int `json:"pageNo" validate:"required"`   // 页号
	PageSize int `json:"pageSize" validate:"required"` // 单页大小
}
