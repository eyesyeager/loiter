package balancer

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
)

/**
 * 负载均衡算法——顺序轮询
 * @auth eyesYeager
 * @date 2024/1/8 11:04
 */

// pollingCursorByAppMap 游标 by AppHost
// 我目前没有设计废弃App清除逻辑，系统本身就是为个人站长设计的，所以我并不认为这样会导致内存溢出
// 同理，我抛弃了并发安全性，追求更高的性能
var pollingCursorByAppMap = make(map[string]int)

// pollingBalancer 顺序轮询负载策略
func pollingBalancer(host string) (error, string) {
	// 获取应用实例
	if _, ok := container.ServerByAppMap[host]; !ok {
		errMsg := fmt.Sprintf("polling load balancing policy execution failed, the application whose host is %s is not registered", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), ""
	}
	// 部分情况下无需执行负载策略
	serverWeight := container.ServerByAppMap[host]
	serverWeightLen := len(serverWeight)
	if serverWeightLen == 0 {
		return errors.New("polling load balancing policy execution failed, payload container cannot be empty"), ""
	}
	if serverWeightLen == 1 {
		return nil, serverWeight[0].Server
	}
	// 处理参数，执行顺序轮询策略
	if _, ok := pollingCursorByAppMap[host]; !ok {
		pollingCursorByAppMap[host] = 0
	}
	url, newCur := pollingActuator(serverWeight, pollingCursorByAppMap[host])
	pollingCursorByAppMap[host] = newCur
	return nil, url
}

// pollingActuator 顺序轮询负载策略执行器
func pollingActuator(serverWeight []container.ServerWeight, cursor int) (string, int) {
	// 为了避免指针越界，此处应该先加一取余，再访问切片
	cursor = (cursor + 1) % len(serverWeight)
	return serverWeight[cursor].Server, cursor
}
