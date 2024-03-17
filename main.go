package main

import (
	"loiter/bootstrap"
	"loiter/kernel"
)

/**
 * @title				loiter(闲游)
 * @version				0.1
 * @date				2023/4/9 20:08
 * @github				https://github.com/YuJiZhao/loiter
 * @author				eyesYeager(耶瞳)
 * @contact.url			http://space.eyesspace.top
 * @contact.email		eyesyeager@163.com
 * @license.name		Apache 2.0
 * @license.url			http://www.apache.org/licenses/LICENSE-2.0.html
 */
func main() {
	// 启动基础服务
	bootstrap.Start()

	// 启动网关
	kernel.Start()

	// 处理程序关闭事项
	defer destruction()
}

// destruction 析构方法，统一处理程序关闭事项
func destruction() {
	// 关闭程序环境
	bootstrap.End()
}
