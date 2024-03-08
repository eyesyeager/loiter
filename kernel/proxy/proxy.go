package proxy

import (
	"fmt"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/balancer"
	"loiter/utils"
	"net/http"
	"net/http/httputil"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 路由
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		host := req.Host

		// 请求预处理
		if err, allow := pre(w, req, host); !allow {
			if err == nil {
				post(w, req, nil, host, constants.PostEntrance.Reject, "")
			} else {
				post(w, req, nil, host, constants.PostEntrance.Error, err.Error())
			}
			return
		}

		// 获取代理信息
		err, targetUrl := balancer.Entry(req, host)
		if err != nil {
			errMsg := fmt.Sprintf("the load balancing policy execution failed and the proxy could not be used. Error message: %s", err.Error())
			statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.BadGateway, errMsg)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			return
		}

		// 创建代理
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		// 响应处理
		proxy.ModifyResponse = func(resp *http.Response) error {
			post(w, req, resp, host, constants.PostEntrance.Post, "")
			return nil
		}

		// 执行反向代理
		proxy.ServeHTTP(w, req)
	})
}
