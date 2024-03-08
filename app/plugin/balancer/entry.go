package balancer

import (
	"gorm.io/gorm"
	"loiter/model/entity"
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

// IBalancerConfigList 负载均衡策略切片
var IBalancerConfigList []entity.Balancer

// Register 注册负载均衡策略
func Register() {
	// 注册到config中，帮助完成数据初始化
	IBalancerConfigList = []entity.Balancer{
		{
			Model:   gorm.Model{ID: 1},
			Code:    random,
			Name:    "随机负载",
			Remarks: "随机选定实例，性能较高。",
		},
		{
			Model:   gorm.Model{ID: 2},
			Code:    polling,
			Name:    "顺序轮询",
			Remarks: "依次选定各个实例，过程中会加锁，性能一般。",
		},
		{
			Model:   gorm.Model{ID: 3},
			Code:    pollingWeighted,
			Name:    "加权轮询",
			Remarks: "根据实例权重选择对应实例，过程中会加锁，性能一般",
		},
	}

	// 注册到Map中，帮助完成网关流程
	IBalancerByNameMap[random] = randomBalancer
	IBalancerByNameMap[polling] = pollingBalancer
	IBalancerByNameMap[pollingWeighted] = pollingWeightedBalancer
}
