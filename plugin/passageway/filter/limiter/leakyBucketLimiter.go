package limiter

import (
	"sync"
	"time"
)

/**
 * 漏桶限流器
 * @auth eyesYeager
 * @date 2024/1/15 10:07
 */

// LeakyBucketLimiter 漏桶限流器
type LeakyBucketLimiter struct {
	mutex     sync.Mutex // 锁
	rate      int        // 令牌产生速率(每秒)
	inventory int        // 当前令牌数
	lastTime  int64      // 上一次滴水的时间戳
}

// TryAcquire 限流检测
func (l *LeakyBucketLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	nowTime := time.Now().Unix()
	interval := int(nowTime - l.lastTime)
	// 滴水并补充令牌
	if interval > 0 {
		l.lastTime = nowTime
		l.inventory = l.rate
	}
	// 消耗令牌
	if l.inventory == 0 {
		return false
	}
	l.inventory--
	return true
}

// LeakyBucketParameter 初始化参数
type LeakyBucketParameter struct {
	Rate int `json:"rate"`
}

// NewLeakyBucketLimiter 创建漏桶限流器
func NewLeakyBucketLimiter(parameter LeakyBucketParameter) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		rate:      parameter.Rate,
		inventory: parameter.Rate,
		lastTime:  time.Now().Unix(),
	}
}
