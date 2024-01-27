package loadbalancer

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/container"
)

/**
 * 负载均衡算法——加权轮询
 * @link https://segmentfault.com/a/1190000018646548
 * @auth eyesYeager
 * @date 2024/1/8 11:05
 */

// weightNodeByAppMap 权重节点 by AppHost
// 我目前没有设计废弃App清除逻辑，系统本身就是为个人站长设计的，所以我并不认为这样会导致内存溢出
// 同理，我抛弃了并发安全性，追求更高的性能
var weightNodeByAppMap = make(map[string][]weightNode)

type weightNode struct {
	server        string // 服务地址
	weight        int    // 权重值
	currentWeight int    // 当前权重
}

// pollingWeightedBalancer 加权轮询负载策略
func pollingWeightedBalancer(host string) (error, string) {
	// 获取应用实例
	if _, ok := container.ServerByAppMap[host]; !ok {
		errMsg := fmt.Sprintf("pollingWeighted load balancing policy execution failed, the application whose host is %s is not registered", host)
		global.GatewayLogger.Error(errMsg)
		return errors.New(errMsg), ""
	}
	// 部分情况下无需执行具体策略
	serverWeight := container.ServerByAppMap[host]
	serverWeightLen := len(serverWeight)
	if serverWeightLen == 0 {
		return errors.New("pollingWeighted load balancing policy execution failed, payload container cannot be empty"), ""
	}
	if serverWeightLen == 1 {
		return nil, serverWeight[0].Server
	}
	// 处理参数，执行加权轮询策略
	if _, ok := weightNodeByAppMap[host]; !ok {
		weightNodeSlice := make([]weightNode, serverWeightLen)
		for index, item := range serverWeight {
			weightNodeSlice[index] = weightNode{
				server:        item.Server,
				weight:        int(item.Weight),
				currentWeight: int(item.Weight),
			}
		}
		weightNodeByAppMap[host] = weightNodeSlice
	}
	return nil, pollingWeightedActuator(weightNodeByAppMap[host])
}

// pollingWeightedActuator 加权轮询负载策略执行期
func pollingWeightedActuator(weightNodeSlice []weightNode) string {
	// 获取总权重与最大权重
	totalWeight := 0
	maxNodeIndex := 0
	for index, item := range weightNodeSlice {
		totalWeight += item.weight
		if item.currentWeight > weightNodeSlice[maxNodeIndex].currentWeight {
			maxNodeIndex = index
		}
	}
	// 更新当前权重
	weightNodeSlice[maxNodeIndex].currentWeight -= totalWeight
	for index := range weightNodeSlice {
		weightNodeSlice[index].currentWeight += weightNodeSlice[index].weight
	}
	return weightNodeSlice[maxNodeIndex].server
}
