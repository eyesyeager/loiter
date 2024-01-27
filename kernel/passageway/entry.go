package passageway

import (
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/plugin/passageway"
	"net/http"
)

/**
 * 通道入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// Entry 进入请求通道
func Entry(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	passagewayNameSlice := container.PassagewayByAppMap[host]
	// 未配置通道则直接放行
	if passagewayNameSlice == nil {
		return nil, true
	}
	// 配置通道则有序进入通道
	for _, name := range passagewayNameSlice {
		if _, ok := passageway.PassageByNameMap[name]; !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no passageway named %s, please deal with it as soon as possible!", name))
			continue
		}
		err, success := passageway.PassageByNameMap[name](w, r, host)
		if err != nil {
			return err, false
		}
		if !success {
			return nil, false
		}
	}
	return nil, true
}
