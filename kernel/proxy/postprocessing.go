package proxy

import (
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/aid"
	"loiter/plugin/store"
	"net/http"
)

/**
 * 代理后置处理
 * @auth eyesYeager
 * @date 2024/1/29 15:32
 */

const (
	preEntrance      = "pre"
	balancerEntrance = "balancer"
	postEntrance     = "post"
)

// 判断是否是异常来源
func isException(entrance string) bool {
	return entrance == preEntrance || entrance == balancerEntrance
}

// post 后置处理总入口
func post(w http.ResponseWriter, req *http.Request, resp *http.Response, host string, entrance string) {
	if entrance == postEntrance {
		entryAid(w, req, resp, host)
	}
	if entrance == preEntrance || entrance == balancerEntrance {
		entryException(w, req, resp, host)
	}
	entryFinal(w, req, resp, host)
	destroyStore(req)
}

// entryAid 进入响应处理器
func entryAid(w http.ResponseWriter, req *http.Request, resp *http.Response, host string) {
	if err := aid.Entry(w, req, resp, host); err != nil {
		errMsg := fmt.Sprintf("aid execution failed. Error message: %s", err.Error())
		statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, errMsg)
		helper.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(errMsg)
	}
}

// entryException 进入异常处理器
func entryException(w http.ResponseWriter, req *http.Request, resp *http.Response, host string) {

}

// entryFinal 进入最终处理器
func entryFinal(w http.ResponseWriter, req *http.Request, resp *http.Response, host string) {

}

// destroyStore 清除状态
func destroyStore(req *http.Request) {
	store.DestroyAll(req)
}
