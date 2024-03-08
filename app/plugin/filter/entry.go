package filter

import (
	"loiter/constants"
	"loiter/model/entity"
	"net/http"
)

/**
 * 过滤器入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

const (
	Limiter  = "limiter"
	NameList = "nameList"
)

// IFilter 请求过滤器方法类型
type IFilter func(http.ResponseWriter, *http.Request, string) (error, bool)

// IFilterByNameMap 过滤器 by 过滤器名
var IFilterByNameMap = make(map[string]IFilter)

// IFilterConfigList 过滤器策略切片
var IFilterConfigList []entity.Processor

// Register 注册过滤器
func Register() {
	// 注册到config中，帮助完成数据初始化
	IFilterConfigList = []entity.Processor{
		{
			Code:    Limiter,
			Name:    "限流器",
			Genre:   constants.Processor.Filter.Code,
			Remarks: "提供多种策略进行限流",
		},
		{
			Code:    NameList,
			Name:    "黑白名单",
			Genre:   constants.Processor.Filter.Code,
			Remarks: "本插件利用布隆过滤器进行初筛，因此尽量避免频繁更新IP",
		},
	}

	// 注册到Map中，帮助完成网关流程
	IFilterByNameMap[Limiter] = LimiterFilter
	IFilterByNameMap[NameList] = NameListFilter
}
