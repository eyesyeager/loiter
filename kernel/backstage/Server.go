package backstage

import (
	"fmt"
	"net/http"
	"zliway/global"
)

/**
 * 后台管理服务
 * @author eyesYeager
 * @date 2023/4/11 15:45
 */

// Server 启动后台管理服务
func Server() {

	fmt.Println("start running backstage service, service port:" + global.Config.App.BackstagePort)
	if err := http.ListenAndServe(":"+global.Config.App.BackstagePort, nil); err != nil {
		panic(fmt.Errorf("failed to execute http.ListenAndServe(:%s): %s", global.Config.App.BackstagePort, err))
	}
}
