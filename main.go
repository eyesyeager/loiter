package main

import (
	"net/http"
	"zliway/test"
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
	// 开启测试web服务器
	test.Web()
	// 开启网关服务
	_ = http.ListenAndServe(":9500", nil)
}
