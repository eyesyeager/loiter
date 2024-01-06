package pipeline

import "net/http"

/**
 * 管道入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

// FrontEntry 管道前置入口
func FrontEntry(host string, w http.ResponseWriter, r *http.Request) {

}

// RearEntry 管道后置入口
func RearEntry(host string, w http.ResponseWriter, r *http.Request) {

}
