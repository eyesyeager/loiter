package aid

import (
	"loiter/kernel/model/entity"
	"net/http"
)

/**
 * 响应处理器入口
 * @auth eyesYeager
 * @date 2024/1/25 17:17
 */

// IAid 响应处理器接口
type IAid func(http.ResponseWriter, *http.Request, *http.Response, string) error

// IAidByNameMap 过滤器方法列表 by 过滤器名
var IAidByNameMap = make(map[string]IAid)

// IAidConfigSlice 响应处理器切片
var IAidConfigSlice []entity.Aid

// Register 注册插件
func Register() {
}
