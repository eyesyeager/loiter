package store

import (
	"errors"
	"fmt"
	"net/http"
)

/**
 * 状态管理
 * 用于上游拦截器向下游拦截器传递数据，类似于数据总线，由于是通过请求头传输，所以需要统一管理，避免出现冲突
 * @auth eyesYeager
 * @date 2024/1/26 10:50
 */

const (
	RequestId        = "Loiter-Store-Request-Id"         // 请求Id
	RequestBeginTime = "Loiter-Store-Request-Begin-Time" // 请求初始时间（时间戳，ms）
)

var keyMap = make(map[string]struct{})

// init 初始化函数
func init() {
	keyMap[RequestId] = struct{}{}
	keyMap[RequestBeginTime] = struct{}{}
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
