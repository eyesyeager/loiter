package balancer

import (
	"errors"
	"fmt"
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

// loadBalancer 负载均衡策略方法类型
type loadBalancer func(string) (error, string)

// balancerByNameMap 负载均衡策略 by 策略名
var balancerByNameMap = make(map[string]loadBalancer)

// InitLoadBalancer 加载负载均衡器
func InitLoadBalancer() {
	balancerByNameMap = make(map[string]loadBalancer)
	balancerByNameMap["random"] = randomBalancer
	balancerByNameMap["polling"] = pollingBalancer
	balancerByNameMap["pollingWeighted"] = pollingWeightedBalancer
}

// Entry 进入负载均衡
func Entry(r *http.Request, host string) (error, *url.URL) {
	// 获取host对应的负载均衡策略
	if _, ok := container.BalanceByAppMap[host]; !ok {
		errMsg := fmt.Sprintf("the application whose host is %s does not have a load balancing policy configured", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), nil
	}
	strategy := container.BalanceByAppMap[host]

	// 执行负载策略
	err, targetUrl := balancerByNameMap[strategy](host)
	if err != nil {
		errMsg := fmt.Sprintf("the application whose host is %s fails to execute load balancing policy, error: %s", targetUrl, err.Error())
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), nil
	}

	// 构建代理
	return nil, &url.URL{
		Scheme: "http", // 暂时只实现http代理
		Host:   targetUrl,
		Path:   r.RequestURI,
	}
}
