package exception

import (
	"fmt"
	"loiter/app/plugin/exception"
	"loiter/global"
	"loiter/kernel/container"
	"net/http"
)

/**
 * 异常处理器
 * @auth eyesYeager
 * @date 2024/2/5 14:10
 */

func Entry(w http.ResponseWriter, req *http.Request, host string, errInfo string) error {
	exceptionNameList := container.ExceptionByAppMap[host]
	// 未配置异常处理器则直接放行
	if exceptionNameList == nil {
		return nil
	}
	// 配置异常处理器则有序执行
	for _, name := range exceptionNameList {
		iException, ok := exception.IExceptionByNameMap[name]
		if !ok {
			global.GatewayLogger.Warn(fmt.Sprintf("there is no Exception named %s, please deal with it as soon as possible!", name))
			continue
		}
		go iException(w, req, host, errInfo)
	}
	return nil
}
