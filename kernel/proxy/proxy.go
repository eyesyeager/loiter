package proxy

import (
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/kernel/balancer"
	"loiter/kernel/helper"
	"loiter/kernel/passageway"
	"net/http"
	"net/http/httputil"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 初始化配置
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host := r.Host

		// 进入通道
		err, success := passageway.Entry(w, r, host)
		if err != nil {
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, constant.ResponseNotice.Empty)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(fmt.Sprintf("channel filtering execution failed. Error message: %s", err.Error()))
			return
		}
		if !success {
			return
		}

		// 获取代理信息
		err, targetUrl := balancer.Entry(r, host)
		if err != nil {
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, constant.ResponseNotice.Empty)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(fmt.Sprintf("the load balancing policy execution failed and the proxy could not be used. Error message: %s", err.Error()))
			return
		}

		// 创建代理
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		// 执行反向代理
		proxy.ServeHTTP(w, r)
	})
}
