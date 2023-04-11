package main

import (
	"zliway/bootstrap"
	"zliway/global"
)

/**
 * @title				zliway(闲游)
 * @version				0.1
 * @date				2023/4/9 20:08
 * @author				eyesYeager(耶瞳)
 * @contact.url			http://space.eyesspace.top
 * @contact.email		eyesyeager@163.com
 * @license.name		Apache 2.0
 * @license.url			http://www.apache.org/licenses/LICENSE-2.0.html
 */
func main() {
	// 初始化程序
	bootstrap.InitializeConfig() // 读取配置文件
	bootstrap.InitializeLog()    // 初始化日志
	bootstrap.InitializeMDB()    // 初始化持久层——MySQL
	bootstrap.InitializeRDB()    // 初始化持久层——Redis

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.MDB != nil {
			db, _ := global.MDB.DB()
			_ = db.Close()
		}
	}()

	// 开启测试web服务器
	//test.Web()
	// 开启网关服务
	//_ = http.ListenAndServe(":"+global.Config.App.Port, nil)
}
