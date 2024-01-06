package constant

import (
	"fmt"
	"loiter/kernel/backstage/model/structure"
)

/**
 * 日志常量
 * @auth eyesYeager
 * @date 2024/1/5 09:45
 */

var LogUniversal = logUniversal{
	DoRegister: logStructure{
		Title:   "开通新账号",
		Content: "账号名：%s；邮箱：%s；备注：%s",
	},
	AddApp: logStructure{
		Title:   "注册应用",
		Content: "应用名：%s；Host：%s；备注：%s",
	},
	AddServer: logStructure{
		Title:   "注册应用实例",
		Content: "应用名：%s；实例名：%s；实例地址：%s；备注：%s",
	},
}

type logUniversal struct {
	DoRegister logStructure
	AddApp     logStructure
	AddServer  logStructure
}

type logStructure struct {
	Title   string
	Content string
}

// BuildUniversalLog 构建通用日志结构
func BuildUniversalLog(log logStructure, params ...any) structure.LogUniversalStruct {
	return structure.LogUniversalStruct{
		Title:   log.Title,
		Content: fmt.Sprintf(log.Content, params...),
	}
}
