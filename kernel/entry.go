package kernel

import (
	"fmt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage"
	"net/http"
)

/**
 * 网关启动器
 * @author eyesYeager
 * @date 2023/9/25 16:50
 */

// Start 启动网关服务
func Start() {
	// 加载注册信息

	// 执行代理配置

	// 启动后台web服务
	backstage.Start()

	// 启动网关服务
	global.AppLogger.Info("start running gateway service, service port:" + config.Program.GateWayPort)
	if err := http.ListenAndServe(":"+config.Program.GateWayPort, nil); err != nil {
		panic(fmt.Errorf("failed to execute http.ListenAndServe(:%s): %s", config.Program.GateWayPort, err))
	}
}
