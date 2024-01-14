package filter

import "loiter/kernel/passageway/genre"

/**
 * 过滤器集合
 * @auth eyesYeager
 * @date 2024/1/9 18:07
 */

// InitFilter 初始化过滤器
func InitFilter() map[string]genre.Aisle {
	filterMap := make(map[string]genre.Aisle)
	filterMap["limiter"] = LimiterFilter
	filterMap["reviewList"] = ReviewListFilter
	return filterMap
}
