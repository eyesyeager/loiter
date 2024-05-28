package final

import (
	"fmt"
	"loiter/app/plugin/final"
	"loiter/global"
	"loiter/kernel/container"
	"net/http"
)

/**
 * 最终处理器
 * @auth eyesYeager
 * @date 2024/2/5 14:10
 */

func Entry(w http.ResponseWriter, req *http.Request, resp *http.Response, host string, entrance string, errInfo string) error {
	finalNameList := container.FinalByAppMap[host]
	// 未配置最终处理器则直接放行
	if finalNameList == nil {
		return nil
	}
	// 配置了最终处理器则有序进入
	for _, name := range finalNameList {
		iFinal, ok := final.IFinalByNameMap[name]
		if !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no Final named %s, please deal with it as soon as possible!", name))
			continue
		}
		err := iFinal(w, req, resp, host, entrance, errInfo)
		if err != nil {
			return err
		}
	}
	return nil
}
