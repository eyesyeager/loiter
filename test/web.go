package test

import (
	"net/http"
)

/**
 * 测试入口
 * @author eyesYeager
 * @date 2023/4/9 20:45
 */

type webHandlerA struct {
}

func (webHandlerA) ServeHTTP(write http.ResponseWriter, r *http.Request) {
	_, _ = write.Write([]byte("1"))
}

func StartWebA() {
	_ = http.ListenAndServe(":9501", webHandlerA{})
}

type webHandlerB struct {
}

func (webHandlerB) ServeHTTP(write http.ResponseWriter, _ *http.Request) {
	_, _ = write.Write([]byte("2"))
}

func StartWebB() {
	_ = http.ListenAndServe(":9502", webHandlerB{})
}

type webHandlerC struct {
}

func (webHandlerC) ServeHTTP(write http.ResponseWriter, _ *http.Request) {
	_, _ = write.Write([]byte("3"))
}

func StartWebC() {
	_ = http.ListenAndServe(":9503", webHandlerC{})
}
