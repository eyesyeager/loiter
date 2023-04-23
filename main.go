package main

import (
	"zliway/bootstrap"
	"zliway/global"
	"zliway/kernel"
	"zliway/test"
)

/**
 * @title				zliway(闲游)
 * @version				0.1
 * @date				2023/4/9 20:08
 * @github				https://github.com/YuJiZhao/zliway
 * @author				eyesYeager(耶瞳)
 * @contact.url			http://space.eyesspace.top
 * @contact.email		eyesyeager@163.com
 * @license.name		Apache 2.0
 * @license.url			http://www.apache.org/licenses/LICENSE-2.0.html
 */
func main() {
	// 启动基础服务
	bootstrap.Start()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.MDB != nil {
			db, _ := global.MDB.DB()
			_ = db.Close()
		}
	}()

	// 启动测试服务
	test.Web()

	// 启动网关服务
	kernel.Start()
}
