package limiter

import (
	"sync"
	"time"
)

/**
 * 令牌桶限流器
 * @auth eyesYeager
 * @date 2024/1/15 10:08
 */

// TokenBucketLimiter 令牌桶限流器
type TokenBucketLimiter struct {
	mutex     sync.Mutex // 锁
	rate      int        // 令牌产生速率(每秒)
	bucket    int        // 桶的容量
	inventory int        // 桶当前令牌数
	lastTime  int64      // 上一次补充令牌的时间戳
}

// TryAcquire 限流检测
func (l *TokenBucketLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	nowTime := time.Now().Unix()
	interval := int(nowTime - l.lastTime)
	// 更新lastTime并尝试补充令牌桶
	if interval > 0 {
		l.lastTime = nowTime
		total := l.inventory + interval*l.rate
		if total > l.bucket {
			l.inventory = l.bucket
		} else {
			l.inventory = total
		}
	}
	// 消耗令牌
	if l.inventory == 0 {
		return false
	}
	l.inventory--
	return true
}

// TokenBucketParameter 初始化参数
type TokenBucketParameter struct {
	Rate   int `json:"rate"`
	Bucket int `json:"bucket"`
}

// NewTokenBucketLimiter 创建令牌桶限流器
func NewTokenBucketLimiter(parameter TokenBucketParameter) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		rate:      parameter.Rate,
		bucket:    parameter.Bucket,
		inventory: parameter.Bucket,
		lastTime:  time.Now().Unix(),
	}
}
