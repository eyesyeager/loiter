package limiter

import (
	"errors"
	"sync"
	"time"
)

/**
 * 滑动窗口限流器
 * @auth eyesYeager
 * @date 2024/1/15 10:06
 */

// SlidingWinLimiter 滑动窗口限流器
type SlidingWinLimiter struct {
	mutex        sync.Mutex    // 避免并发问题
	limit        int           // 窗口请求上限
	window       int64         // 窗口时间大小(ms)
	smallWindow  int64         // 小窗口时间大小(ms)
	smallWindows int64         // 小窗口数量
	counters     map[int64]int // 小窗口计数器
}

// TryAcquire 限流检测
func (l *SlidingWinLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 获取当前小窗口值
	currentSmallWindow := time.Now().UnixMilli() / l.smallWindow * l.smallWindow
	// 获取起始小窗口值
	startSmallWindow := currentSmallWindow - l.smallWindow*(l.smallWindows-1)

	// 计算当前窗口的请求总数
	var count int
	for smallWindow, counter := range l.counters {
		if smallWindow < startSmallWindow {
			delete(l.counters, smallWindow)
		} else {
			count += counter
		}
	}

	// 若到达窗口请求上限，请求失败
	if count >= l.limit {
		return false
	}
	// 若没到窗口请求上限，当前小窗口计数器+1，请求成功
	l.counters[currentSmallWindow]++
	return true
}

// SlidingWinParameter 初始化参数
type SlidingWinParameter struct {
	Limit       int   `json:"limit"`
	Window      int64 `json:"window"`
	SmallWindow int64 `json:"smallWindow"`
}

// NewSlidingWinLimiter 创建滑动窗口限流器
func NewSlidingWinLimiter(parameter SlidingWinParameter) (*SlidingWinLimiter, error) {
	// 窗口时间必须能够被小窗口时间整除
	if parameter.Window%parameter.SmallWindow != 0 {
		return nil, errors.New("window cannot be split by integers")
	}

	return &SlidingWinLimiter{
		limit:        parameter.Limit,
		window:       parameter.Window,
		smallWindow:  parameter.SmallWindow,
		smallWindows: parameter.Window / parameter.SmallWindow,
		counters:     make(map[int64]int),
	}, nil
}
