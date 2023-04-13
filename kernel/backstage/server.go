package backstage

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"zliway/global"
	"zliway/kernel/backstage/router"
)

/**
 * 后台管理服务
 * @author eyesYeager
 * @date 2023/4/11 15:45
 */

// Server zliway后台服务
func Server() {
	// 初始化路由
	routerRoot := httprouter.New()
	router.InitRouter(routerRoot)
	// 启动服务
	fmt.Println("start running backstage service, service port:" + global.Config.App.BackstagePort)
	if err := http.ListenAndServe(":"+global.Config.App.BackstagePort, routerRoot); err != nil {
		panic(fmt.Errorf("failed to execute http.ListenAndServe(:%s): %s", global.Config.App.BackstagePort, err))
	}
}
