package limiter

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/kernel/model/entity"
)

/**
 * 限流器整合
 * @auth eyesYeager
 * @date 2024/1/22 19:28
 */

// ILimiter 限流器接口
type ILimiter interface {
	TryAcquire() bool
}

// NewLimiterFilter 整合限流器创建方法，供容器初始化/刷新使用
func NewLimiterFilter(name string, param string) (error, ILimiter) {
	if name == LimiterConfig.FixedWinLimiter.Name { // 固定窗口限流
		var parameter FixedWinParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the FixedWinLimiter are not in this format: %s", param)), nil
		}
		return nil, NewFixedWinLimiter(parameter)
	} else if name == LimiterConfig.SlidingWinLimiter.Name {
		//NewSlidingWinLimiter()
		return nil, nil
	} else if name == LimiterConfig.LeakyBucketLimiter.Name {
		var parameter LeakyBucketParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the LeakyBucketParameter are not in this format: %s", param)), nil
		}
		return nil, NewLeakyBucketLimiter(parameter)
	} else if name == LimiterConfig.TokenBucketLimiter.Name {
		var parameter TokenBucketParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the TokenBucketLimiter are not in this format: %s", param)), nil
		}
		return nil, NewTokenBucketLimiter(parameter)
	} else {
		return errors.New(fmt.Sprintf("there is no current limiter named %s", name)), nil
	}
}

/***************************************
 *            限流器实体数据
 ***************************************/

var LimiterConfig = limiterConfig{
	FixedWinLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 1},
		Name:      "FixedWinLimiter",
		Parameter: "{\n    \"limit\": 1,  // 数字类型，窗口请求次数上限\n    \"window\": 1  // 数字类型，窗口时间大小(单位为秒)\n}",
		Remarks:   "固定窗口限流算法",
	},
	SlidingWinLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 2},
		Name:      "SlidingWinLimiter",
		Parameter: "",
		Remarks:   "滑动窗口限流算法",
	},
	LeakyBucketLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 3},
		Name:      "LeakyBucketLimiter",
		Parameter: "{\n  \"rate\": 1   // 数字类型，令牌产生速率(每秒)\n}",
		Remarks:   "漏桶限流算法",
	},
	TokenBucketLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 4},
		Name:      "TokenBucketLimiter",
		Parameter: "{\n  \"rate\": 1,   // 数字类型，令牌产生速率(每秒)\n  \"bucket\": 1  // 令牌桶容量\n}",
		Remarks:   "令牌桶限流算法",
	},
}

type limiterConfig struct {
	FixedWinLimiter    entity.Limiter
	SlidingWinLimiter  entity.Limiter
	LeakyBucketLimiter entity.Limiter
	TokenBucketLimiter entity.Limiter
}
