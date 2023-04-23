package test

import "net/http"

/**
 * 测试服务器
 * @author eyesYeager
 * @date 2023/4/10 11:24
 */

type webHandlerA struct {
}

func (webHandlerA) ServeHTTP(write http.ResponseWriter, _ *http.Request) {
	_, _ = write.Write([]byte("test webA"))
}

func StartWebA() {
	_ = http.ListenAndServe(":9501", webHandlerA{})
}

type webHandlerB struct {
}

func (webHandlerB) ServeHTTP(write http.ResponseWriter, _ *http.Request) {
	_, _ = write.Write([]byte("test webB"))
}

func StartWebB() {
	_ = http.ListenAndServe(":9502", webHandlerB{})
}

type webHandlerC struct {
}

func (webHandlerC) ServeHTTP(write http.ResponseWriter, _ *http.Request) {
	_, _ = write.Write([]byte("test webC"))
}

func StartWebC() {
	_ = http.ListenAndServe(":9503", webHandlerC{})
}
