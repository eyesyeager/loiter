package proxy

import (
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/balancer"
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

		// 请求预处理
		if allow := pre(w, req, host); !allow {
			post(w, req, nil, host, preEntrance)
			return
		}

		// 获取代理信息
		err, targetUrl := balancer.Entry(req, host)
		if err != nil {
			errMsg := fmt.Sprintf("the load balancing policy execution failed and the proxy could not be used. Error message: %s", err.Error())
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, errMsg)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			post(w, req, nil, host, balancerEntrance)
			return
		}

		// 创建代理
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		// 响应处理
		proxy.ModifyResponse = func(resp *http.Response) error {
			post(w, req, resp, host, postEntrance)
			return nil
		}

		// 执行反向代理
		proxy.ServeHTTP(w, req)
	})
}
