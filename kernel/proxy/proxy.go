package proxy

import (
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/aid"
	"loiter/kernel/balancer"
	"loiter/kernel/passageway"
	"loiter/kernel/store"
	"net/http"
	"net/http/httputil"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 初始化配置
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		host := req.Host

		// 进入通道
		err, success := passageway.Entry(w, req, host)
		if err != nil {
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, constant.ResponseNotice.Empty)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(fmt.Sprintf("passageway execution failed. Error message: %s", err.Error()))
			return
		}
		if !success {
			return
		}

		// 获取代理信息
		err, targetUrl := balancer.Entry(req, host)
		if err != nil {
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, constant.ResponseNotice.Empty)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(fmt.Sprintf("the load balancing policy execution failed and the proxy could not be used. Error message: %s", err.Error()))
			return
		}

		// 创建代理
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		// 响应处理
		proxy.ModifyResponse = func(resp *http.Response) error {
			// 进入响应总线
			if err = aid.Entry(w, req, resp, host); err != nil {
				statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, constant.ResponseNotice.Empty)
				helper.Response(w, statusCode, contentType, content)
				global.GatewayLogger.Warn(fmt.Sprintf("aid execution failed. Error message: %s", err.Error()))
			}
			// 销毁状态
			store.DestroyAll(req)
			return nil
		}

		// 执行反向代理
		proxy.ServeHTTP(w, req)
	})
}
