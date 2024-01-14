package constant

import (
	"fmt"
	"loiter/kernel/model/structure"
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
	UpdateAppBalance: logStructure{
		Title:   "更新应用负载均衡策略",
		Content: "应用名：%s；原负载策略：%s；更新后负载策略：%s",
	},
	RefreshAllContainer: logStructure{
		Title:   "刷新指定应用下的所有容器",
		Content: "应用名：%s；",
	},
	RefreshAppServerContainer: logStructure{
		Title:   "刷新指定应用下的应用实例容器",
		Content: "应用名：%s；",
	},
	RefreshBalanceContainer: logStructure{
		Title:   "刷新指定应用下的负载均衡容器",
		Content: "应用名：%s；",
	},
	UpdateAppPassageway: logStructure{
		Title:   "更新应用通道配置",
		Content: "应用名：%s；原通道配置：%s；更新后通道配置：%s",
	},
}

type logUniversal struct {
	DoRegister                logStructure
	AddApp                    logStructure
	AddServer                 logStructure
	UpdateAppBalance          logStructure
	RefreshAllContainer       logStructure
	RefreshAppServerContainer logStructure
	RefreshBalanceContainer   logStructure
	UpdateAppPassageway       logStructure
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
