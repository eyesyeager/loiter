package filter

import (
	"errors"
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/container"
	"net/http"
)

/**
 * 限流器
 * @link: https://juejin.cn/post/7056068978862456846
 * @auth eyesYeager
 * @date 2024/1/11 16:46
 */

// LimiterFilter 限流入口方法
func LimiterFilter(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	if limiter, ok := container.LimiterByAppMap[host]; ok {
		success := limiter.TryAcquire()
		if !success {
			errMsg := fmt.Sprintf("the application with host %s is throttled", host)
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.RateLimit, errMsg)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
		}
		return nil, success
	}
	return errors.New(fmt.Sprintf("the application with host %s does not specify a current limiting algorithm", host)), false
}
