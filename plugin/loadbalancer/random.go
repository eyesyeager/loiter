package loadbalancer

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
	"math/rand"
	"time"
)

/**
 * 负载均衡算法——随机负载
 * @auth eyesYeager
 * @date 2024/1/8 11:04
 */

// 随机数生成器
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// randomBalancer 随机负载策略
func randomBalancer(host string) (error, string) {
	// 获取应用实例
	if _, ok := container.ServerByAppMap[host]; !ok {
		errMsg := fmt.Sprintf("random load balancing policy execution failed, the application whose host is %s is not registered", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), ""
	}
	// 部分情况下无需执行具体策略
	serverWeight := container.ServerByAppMap[host]
	serverWeightLen := len(serverWeight)
	if serverWeightLen == 0 {
		return errors.New("random load balancing policy execution failed, payload container cannot be empty"), ""
	}
	if serverWeightLen == 1 {
		return nil, serverWeight[0].Server
	}
	// 执行随机负载策略
	return nil, randomActuator(serverWeight)
}

// randomActuator 随机负载策略执行器
func randomActuator(serverWeight []container.ServerWeight) string {
	serverWeightLen := len(serverWeight)
	return serverWeight[rnd.Intn(serverWeightLen)].Server
}
