package passageway

import (
	"gorm.io/gorm"
	"loiter/kernel/model/entity"
	"loiter/plugin/passageway/filter"
	"loiter/plugin/passageway/pipeline"
	"net/http"
)

/**
 * 通道入口
 * @auth eyesYeager
 * @date 2024/1/5 16:14
 */

const (
	Filter   = "filter"   // 过滤器
	Pipeline = "pipeline" // 管道

	RequestLog = "requestLog"
	Limiter    = "limiter"
	NameList   = "nameList"
)

// IPassageway 请求通道方法类型
type IPassageway func(http.ResponseWriter, *http.Request, string) (error, bool)

// PassageByNameMap 请求通道方法列表 by 通道名
var PassageByNameMap = make(map[string]IPassageway)

// PassageConfigSlice 请求通道策略切片
var PassageConfigSlice []entity.Passageway

// Register 注册请求通道
func Register() {
	// 注册到config中，帮助完成数据初始化
	PassageConfigSlice = []entity.Passageway{
		{
			Model:   gorm.Model{ID: 1},
			Name:    RequestLog,
			Genre:   Pipeline,
			Remarks: "日志管道",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    Limiter,
			Genre:   Filter,
			Remarks: "限流过滤器",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    NameList,
			Genre:   Filter,
			Remarks: "黑白名单过滤器",
		},
	}

	// 注册到Map中，帮助完成网关流程
	PassageByNameMap[RequestLog] = pipeline.RequestLogPipeline
	PassageByNameMap[Limiter] = filter.LimiterFilter
	PassageByNameMap[NameList] = filter.NameListFilter
}
