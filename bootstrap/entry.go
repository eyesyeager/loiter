package bootstrap

import (
	"fmt"
)

/**
 * 基础结构启动入口
 * @author eyesYeager
 * @date 2023/4/11 16:54
 */

func Start() {
	fmt.Println("start starting basic services...")

	// 初始化程序
	initializeConfig() // 读取配置文件
	initializeLog()    // 初始化日志
	initializeMDB()    // 初始化持久层——MySQL
	initializeRDB()    // 初始化持久层——Redis

	fmt.Println("basic service started successfully")
}
