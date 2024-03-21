package constant

import (
	"fmt"
	"loiter/model/structure"
)

/**
 * 日志常量
 * @auth eyesYeager
 * @date 2024/1/5 09:45
 */

var LogUniversal = logUniversal{
	DoLogin: logStructure{
		Title:   "用户登录",
		Content: "IP：%s；浏览器：%s；",
	},
	DoRegister: logStructure{
		Title:   "开通新账号",
		Content: "应用名：%s；邮箱：%s；备注：%s；",
	},
	AddApp: logStructure{
		Title:   "注册应用",
		Content: "注册信息：%s；备注：%s；",
	},
	UpdateApp: logStructure{
		Title:   "更新应用",
		Content: "更新信息：%s；",
	},
	DeleteApp: logStructure{
		Title:   "删除应用",
		Content: "应用名：%s；备注：%s；",
	},
	ActivateApp: logStructure{
		Title:   "变更应用状态",
		Content: "应用名：%s；变更前状态：%s；变更后状态：%s；",
	},
	SaveStaticApp: logStructure{
		Title:   "更新应用静态配置",
		Content: "更新配置：%s；",
	},
	AddServer: logStructure{
		Title:   "注册应用实例",
		Content: "应用名：%s；实例名：%s；实例地址：%s；备注：%s；",
	},
	UpdateAppBalancer: logStructure{
		Title:   "更新应用负载均衡策略",
		Content: "原负载策略：%s；更新后负载策略：%s；",
	},
	RefreshContainer: logStructure{
		Title:   "刷新指定应用下的容器",
		Content: "容器名：%s；",
	},
	SaveAppProcessor: logStructure{
		Title:   "更新指定应用下的处理器",
		Content: "更新配置：%s；",
	},
	UpdateAppLimiter: logStructure{
		Title:   "更新应用限流器配置",
		Content: "原限流器名：%s；原限流器参数：%s；更新限流器名：%s；更新限流器参数：%s；",
	},
	UpdateAppNameList: logStructure{
		Title:   "更新应用黑白名单配置",
		Content: "名单类型：%s；更新类型：%s；",
	},
	AddNameListIp: logStructure{
		Title:   "添加黑白名单ip",
		Content: "应用名：%s；名单类型：%s；添加ip：%s；",
	},
	DeleteNameListIp: logStructure{
		Title:   "删除黑白名单ip",
		Content: "应用名：%s；名单类型：%s；删除ip：%s；",
	},
}

type logUniversal struct {
	DoLogin           logStructure
	DoRegister        logStructure
	AddApp            logStructure
	UpdateApp         logStructure
	DeleteApp         logStructure
	ActivateApp       logStructure
	SaveStaticApp     logStructure
	AddServer         logStructure
	UpdateAppBalancer logStructure
	RefreshContainer  logStructure
	SaveAppProcessor  logStructure
	UpdateAppLimiter  logStructure
	UpdateAppNameList logStructure
	AddNameListIp     logStructure
	DeleteNameListIp  logStructure
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
