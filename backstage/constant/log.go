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
		Content: "应用名：%s；邮箱：%s；备注：%s",
	},
	AddApp: logStructure{
		Title:   "注册应用",
		Content: "应用名：%s；Host：%s；备注：%s",
	},
	AddServer: logStructure{
		Title:   "注册应用实例",
		Content: "应用名：%s；实例名：%s；实例地址：%s；备注：%s",
	},
	UpdateAppBalancer: logStructure{
		Title:   "更新应用负载均衡策略",
		Content: "原负载策略：%s；更新后负载策略：%s",
	},
	RefreshContainer: logStructure{
		Title:   "刷新指定应用下的容器",
		Content: "容器名：%s",
	},
	UpdateAppFilter: logStructure{
		Title:   "更新应用过滤器配置",
		Content: "原过滤器配置：%s；更新后过滤器配置：%s",
	},
	UpdateAppLimiter: logStructure{
		Title:   "更新应用限流器配置",
		Content: "原限流器名：%s；原限流器参数：%s；更新限流器名：%s；更新限流器参数：%s",
	},
	UpdateAppNameList: logStructure{
		Title:   "更新应用黑白名单配置",
		Content: "名单类型：%s；更新类型：%s",
	},
	AddNameListIp: logStructure{
		Title:   "添加黑白名单ip",
		Content: "应用名：%s；名单类型：%s；添加ip：%s",
	},
	DeleteNameListIp: logStructure{
		Title:   "删除黑白名单ip",
		Content: "应用名：%s；名单类型：%s；删除ip：%s",
	},
	UpdateAppAid: logStructure{
		Title:   "更新应用响应处理器",
		Content: "原处理器：%s；更新处理器：%s",
	},
}

type logUniversal struct {
	DoRegister        logStructure
	AddApp            logStructure
	AddServer         logStructure
	UpdateAppBalancer logStructure
	RefreshContainer  logStructure
	UpdateAppFilter   logStructure
	UpdateAppLimiter  logStructure
	UpdateAppNameList logStructure
	AddNameListIp     logStructure
	DeleteNameListIp  logStructure
	UpdateAppAid      logStructure
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
