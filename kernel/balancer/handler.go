package balancer

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"zliway/global"
	"zliway/kernel/backstage/controller/result"
	"zliway/kernel/dispose"
)

/**
 * @author eyesYeager
 * @date 2023/4/23 15:10
 */

// SpecifyURL 获取被代理的URL
func SpecifyURL(r *http.Request) (err error, url *url.URL) {
	model, ok := dispose.AppHolder[r.Host]

	// 判断host是否合法
	if !ok {
		global.Log.Error("Illegal host: " + r.Host)
		return errors.New(result.Results.DefaultFail.Msg), url
	}

	// 根据app模式选择处理方式
	switch model.Pattern {
	case dispose.AppPattern["singleton"]:
		url = singletonServer(model)
	case dispose.AppPattern["micro"]:
		url = microServer(model)
	default:
		errorMsg := "unknown app pattern:" + strconv.Itoa(int(model.Pattern))
		global.Log.Error(errorMsg)
		return errors.New(errorMsg), url
	}

	// 执行path处理
	url.Path = r.RequestURI
	return err, url
}

// 单体app服务
func singletonServer(model dispose.AppHolderModel) *url.URL {
	return executeLoadBalance(model.Balancer, model.Servers)
}

// 微服务app服务
func microServer(model dispose.AppHolderModel) *url.URL {
	return nil
}

// 执行负载均衡
func executeLoadBalance(balancer uint8, servers []dispose.ServerHolderModel) *url.URL {
	server := StrategyMap[balancer](servers)
	return &url.URL{
		Scheme: "http", // TODO: 暂时打算只实现http，其他协议以后再说
		Host:   server,
	}
}
