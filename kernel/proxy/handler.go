package proxy

import (
	"fmt"
	"zliway/kernel/backstage/service"
)

/**
 * 代理处理器
 * @author eyesYeager
 * @date 2023/4/23 15:08
 */

// InitProxy 初始化反向代理配置
func InitProxy() {
	if err := service.AppService.FillAppHolder(); err != nil {
		panic(fmt.Errorf("failed to init proxy: %s", err))
	}
}
