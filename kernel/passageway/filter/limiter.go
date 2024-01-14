package filter

import "net/http"

/**
 * 限流器
 * @link: https://juejin.cn/post/7056068978862456846
 * @auth eyesYeager
 * @date 2024/1/11 16:46
 */

func LimiterFilter(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	return nil, true
}

// fixedWindowLimiter 固定窗口限流
func fixedWindowLimiter() {

}

// slideWindowLimiter 滑动窗口限流
func slideWindowLimiter() {

}

// leakyBucketLimiter 漏桶限流
func leakyBucketLimiter() {

}

// tokenBucketLimiter 令牌桶限流
func tokenBucketLimiter() {

}
