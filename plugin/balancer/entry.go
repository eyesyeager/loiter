package balancer

import (
	"gorm.io/gorm"
	"loiter/kernel/model/entity"
)

/**
 * 负载均衡模块入口类
 * @auth eyesYeager
 * @date 2024/1/22 11:44
 */

const (
	random          = "random"
	polling         = "polling"
	pollingWeighted = "pollingWeighted"
)

// IBalancer 负载均衡策略方法类型
type IBalancer func(string) (error, string)

// IBalancerByNameMap 负载均衡策略 by 策略名
var IBalancerByNameMap = make(map[string]IBalancer)

// IBalancerConfigSlice 负载均衡策略切片
var IBalancerConfigSlice []entity.Balancer

// Register 注册负载均衡策略
func Register() {
	// 注册到config中，帮助完成数据初始化
	IBalancerConfigSlice = []entity.Balancer{
		{
			Model:   gorm.Model{ID: 1},
			Name:    random,
			Remarks: "随机负载",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    polling,
			Remarks: "轮询负载",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    pollingWeighted,
			Remarks: "加权轮询负载",
		},
	}

	// 注册到Map中，帮助完成网关流程
	IBalancerByNameMap[random] = randomBalancer
	IBalancerByNameMap[polling] = pollingBalancer
	IBalancerByNameMap[pollingWeighted] = pollingWeightedBalancer
}
