package aid

import (
	"fmt"
	"loiter/app/plugin/aid"
	"loiter/global"
	"loiter/kernel/container"
	"net/http"
)

/**
 * 响应处理器
 * @auth eyesYeager
 * @date 2024/1/25 17:10
 */

// Entry 进入响应处理器
func Entry(w http.ResponseWriter, req *http.Request, resp *http.Response, host string) error {
	aidNameList := container.AidByAppMap[host]
	// 未配置响应处理器则直接放行
	if aidNameList == nil {
		return nil
	}
	// 配置响应处理器则有序执行
	for _, name := range aidNameList {
		iAid, ok := aid.IAidByNameMap[name]
		if !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no aid named %s, please deal with it as soon as possible!", name))
			continue
		}
		err := iAid(w, req, resp, host)
		if err != nil {
			return err
		}
	}
	return nil
}
