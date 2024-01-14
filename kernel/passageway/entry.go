package passageway

import (
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/kernel/passageway/filter"
	"loiter/kernel/passageway/genre"
	"loiter/kernel/passageway/pipeline"
	"net/http"
)

/**
 * 通道入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// aisleByNameMap 通道方法列表 by 通道名
var aisleByNameMap = make(map[string]genre.Aisle)

// InitAisle 加载负载均衡器
func InitAisle() {
	aisleByNameMap = make(map[string]genre.Aisle)
	for key, value := range filter.InitFilter() {
		aisleByNameMap[key] = value
	}
	for key, value := range pipeline.InitPipeline() {
		aisleByNameMap[key] = value
	}
}

// Entry 进入通道
func Entry(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	passagewayNameSlice := container.PassagewayByAppMap[host]
	if passagewayNameSlice == nil {
		return nil, true
	}
	for _, name := range passagewayNameSlice {
		if _, ok := aisleByNameMap[name]; !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no passageway named %s, please deal with it as soon as possible!", name))
			continue
		}
		err, success := aisleByNameMap[name](w, r, host)
		if err != nil {
			return err, false
		}
		if !success {
			return nil, false
		}
	}
	return nil, true
}
