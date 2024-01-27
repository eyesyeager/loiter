package store

import (
	"errors"
	"fmt"
	"net/http"
)

/**
 * 状态管理
 * 类似于数据总线，由于是通过请求头传输，所以需要统一管理，避免出现冲突
 * @auth eyesYeager
 * @date 2024/1/26 10:50
 */

const (
	RequestLogTime = "Loiter-Store-Request-Log-Time" // 请求日志时间
	RequestLogId   = "Loiter-Store-Request-Log-Id"   // 请求日志id
)

var keyMap = make(map[string]struct{})

// init 初始化函数
func init() {
	keyMap[RequestLogTime] = struct{}{}
	keyMap[RequestLogId] = struct{}{}
}

// SetValue 写入总线数据
func SetValue(r *http.Request, key string, value string) error {
	if _, ok := keyMap[key]; !ok {
		return errors.New(fmt.Sprintf("key '%s' not declared in the store is not allowed to be used", key))
	}
	r.Header.Set(key, value)
	return nil
}

// GetValue 获取总线数据
func GetValue(r *http.Request, key string) string {
	if _, ok := keyMap[key]; !ok {
		return ""
	}
	return r.Header.Get(key)
}

// DestroyAll 销毁所有总线数据
func DestroyAll(r *http.Request) {
	for key := range keyMap {
		r.Header.Del(key)
	}
}
