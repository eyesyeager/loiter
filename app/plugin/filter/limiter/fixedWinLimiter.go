package limiter

import (
	"sync"
	"time"
)

/**
 * 固定窗口限流器
 * @auth eyesYeager
 * @date 2024/1/15 10:05
 */

// FixedWinLimiter 固定窗口限流器
type FixedWinLimiter struct {
	mutex    sync.Mutex // 锁
	limit    int        // 窗口请求上限
	window   int64      // 窗口时间大小(单位为秒)，为了避免每次限流检测时都做类型转换，我将其设置为int64，这样就仅在实例化时做类型转换
	counter  int        // 计数器
	lastTime int64      // 上一次请求的时间戳(单位为秒)
}

// TryAcquire 限流检测
func (l *FixedWinLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	// 获取当前时间
	now := time.Now().Unix()
	// 如果当前窗口失效，计数器清0，开启新的窗口
	if now > l.lastTime+l.window {
		l.counter = 0
		l.lastTime = now
	}
	// 若到达窗口请求上限，请求失败
	if l.counter >= l.limit {
		return false
	}
	// 若没到窗口请求上限，计数器+1，请求成功
	l.counter++
	return true
}

// FixedWinParameter 初始化参数
type FixedWinParameter struct {
	Limit  int `json:"limit"`
	Window int `json:"window"`
}

// NewFixedWinLimiter 创建固定窗口限流器
func NewFixedWinLimiter(parameter FixedWinParameter) *FixedWinLimiter {
	return &FixedWinLimiter{
		limit:    parameter.Limit,
		window:   int64(parameter.Window),
		lastTime: time.Now().Unix(),
	}
}
