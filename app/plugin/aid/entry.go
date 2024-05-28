package aid

import (
	"loiter/model/entity"
	"net/http"
)

/**
 * 响应处理器入口
 * @auth eyesYeager
 * @date 2024/1/25 17:17
 */

// IAid 响应处理器接口
type IAid func(http.ResponseWriter, *http.Request, *http.Response, string) error

// IAidByNameMap 响应处理器方法列表 by 处理器名
var IAidByNameMap = make(map[string]IAid)

// IAidConfigList 响应处理器切片
var IAidConfigList []entity.Processor

// Register 注册插件
func Register() {
}
