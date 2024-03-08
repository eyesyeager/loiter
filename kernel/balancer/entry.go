package balancer

import (
	"errors"
	"fmt"
	"loiter/app/plugin/balancer"
	"loiter/global"
	"loiter/kernel/container"
	"net/http"
	"net/url"
)

/**
 * 负载均衡器入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// Entry 进入负载均衡
func Entry(r *http.Request, host string) (error, *url.URL) {
	// 获取host对应的负载均衡策略
	if _, ok := container.BalancerByAppMap[host]; !ok {
		errMsg := fmt.Sprintf("the application whose host is %s does not have a Balancer", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), nil
	}
	strategy := container.BalancerByAppMap[host]

	// 执行负载策略
	if _, ok := balancer.IBalancerByNameMap[strategy]; !ok {
		errMsg := fmt.Sprintf("the application whose host is %s does not registered in container", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), nil
	}
	err, targetUrl := balancer.IBalancerByNameMap[strategy](host)
	if err != nil {
		errMsg := fmt.Sprintf("the application whose host is %s fails to execute Balancer, error: %s", targetUrl, err.Error())
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), nil
	}

	// 构建代理
	return nil, &url.URL{
		Scheme: "http", // 暂时只实现http代理
		Host:   targetUrl,
	}
}
