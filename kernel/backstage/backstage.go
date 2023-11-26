package backstage

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/router"
	"net/http"
)

/**
 * 后台管理服务
 * @author eyesYeager
 * @date 2023/4/11 15:45
 */

// Start loiter后台服务启动器
func Start() {
	// 初始化路由
	routerRoot := httprouter.New()
	router.InitRouter(routerRoot)

	// 启动服务
	global.AppLogger.Info("start running backstage service, service port:" + config.Program.BackstagePort)
	if err := http.ListenAndServe(":"+config.Program.BackstagePort, routerRoot); err != nil {
		panic(fmt.Errorf("failed to execute http.ListenAndServe(:%s): %s", config.Program.BackstagePort, err))
	}
}

// InitRegister 初始化注册信息
func InitRegister() {

}
