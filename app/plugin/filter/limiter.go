package filter

import (
	"errors"
	"fmt"
	"loiter/app/capability"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/utils"
	"net/http"
)

/**
 * 限流器
 * @link: https://juejin.cn/post/7056068978862456846
 * @auth eyesYeager
 * @date 2024/1/11 16:46
 */

// LimiterFilter 限流入口方法
func LimiterFilter(w http.ResponseWriter, r *http.Request, host string, genre string) (error, bool) {
	if limiter, ok := container.LimiterByAppMap[host]; ok {
		success := limiter.TryAcquire()
		if !success {
			errMsg := fmt.Sprintf("the application with host %s is throttled", host)
			statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.RateLimit, errMsg, genre)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			go capability.NoticeFoundation.SendSiteNotice(host, "限流器触发拦截", errMsg,
				fmt.Sprintf("请求路径：%s；", r.URL.Path))
		}
		return nil, success
	}
	return errors.New(fmt.Sprintf("the application with host %s does not specify a current limiting algorithm", host)), false
}
