package filter

import (
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/plugin/filter"
	"net/http"
)

/**
 * 过滤器入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// Entry 进入请求过滤器
func Entry(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	filterNameSlice := container.FilterByAppMap[host]
	// 未配置过滤器则直接放行
	if filterNameSlice == nil {
		return nil, true
	}
	// 配置过滤器则有序进入
	for _, name := range filterNameSlice {
		iFilter, ok := filter.IFilterByNameMap[name]
		if !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no filter named %s, please deal with it as soon as possible!", name))
			continue
		}
		err, success := iFilter(w, r, host)
		if err != nil {
			return err, false
		}
		if !success {
			return nil, false
		}
	}
	return nil, true
}
