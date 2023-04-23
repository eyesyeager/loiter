package kernel

import (
	"fmt"
	"net/http"
	"zliway/global"
	"zliway/kernel/backstage"
	"zliway/kernel/proxy"
)

/**
 * 网关服务总启动器
 * @author eyesYeager
 * @date 2023/4/11 15:34
 */

// Start 网关服务启动方法
func Start() {
	// 启动后台web服务
	go backstage.Server()

	// 初始化网关配置
	initZliway()

	// 执行代理配置
	proxy.StartProxy()

	// 启动网关服务
	fmt.Println("start running gateway service, service port:" + global.Config.App.Port)
	if err := http.ListenAndServe(":"+global.Config.App.Port, nil); err != nil {
		panic(fmt.Errorf("failed to execute http.ListenAndServe(:%s): %s", global.Config.App.Port, err))
	}
}

// 初始化网关配置
func initZliway() {
	// 初始化代理配置
	proxy.InitProxy()

	// 初始化过滤器配置
}
