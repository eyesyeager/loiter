package limiter

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/model/entity"
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
	if name == LimiterConfig.FixedWinLimiter.Code { // 固定窗口限流
		var parameter FixedWinParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the FixedWinLimiter are not in this format: %s", param)), nil
		}
		return nil, NewFixedWinLimiter(parameter)
	} else if name == LimiterConfig.SlidingWinLimiter.Code {
		var parameter SlidingWinParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the SlidingWinLimiter are not in this format: %s", param)), nil
		}
		limiter, err := NewSlidingWinLimiter(parameter)
		if err != nil {
			return err, nil
		}
		return nil, limiter
	} else if name == LimiterConfig.LeakyBucketLimiter.Code {
		var parameter LeakyBucketParameter
		if err := json.Unmarshal([]byte(param), &parameter); err != nil {
			return errors.New(fmt.Sprintf("The parameters of the LeakyBucketParameter are not in this format: %s", param)), nil
		}
		return nil, NewLeakyBucketLimiter(parameter)
	} else if name == LimiterConfig.TokenBucketLimiter.Code {
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
		Code:      "FixedWinLimiter",
		Name:      "固定窗口",
		Parameter: "{\n    \"limit\": 1,  // 数字类型，窗口请求次数上限\n    \"window\": 1  // 数字类型，窗口时间大小(单位为秒)\n}",
		Remarks:   "",
	},
	SlidingWinLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 2},
		Code:      "SlidingWinLimiter",
		Name:      "滑动窗口",
		Parameter: "{\n  \t\"limit\": 1,  // 窗口请求上限\n\t\"window\": 1000,  // 窗口时间大小(ms)\n\t\"smallWindow\": 200  // 小窗口时间大小(ms)\n}",
		Remarks:   "",
	},
	LeakyBucketLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 3},
		Code:      "LeakyBucketLimiter",
		Name:      "漏桶限流",
		Parameter: "{\n  \"rate\": 1   // 数字类型，令牌产生速率(每秒)\n}",
		Remarks:   "",
	},
	TokenBucketLimiter: entity.Limiter{
		Model:     gorm.Model{ID: 4},
		Code:      "TokenBucketLimiter",
		Name:      "令牌桶限流",
		Parameter: "{\n  \"rate\": 1,   // 数字类型，令牌产生速率(每秒)\n  \"bucket\": 1  // 令牌桶容量\n}",
		Remarks:   "",
	},
}

type limiterConfig struct {
	FixedWinLimiter    entity.Limiter
	SlidingWinLimiter  entity.Limiter
	LeakyBucketLimiter entity.Limiter
	TokenBucketLimiter entity.Limiter
}
