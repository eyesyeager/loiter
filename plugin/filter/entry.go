package filter

import (
	"gorm.io/gorm"
	"loiter/kernel/model/entity"
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

// IFilterConfigSlice 过滤器策略切片
var IFilterConfigSlice []entity.Filter

// Register 注册过滤器
func Register() {
	// 注册到config中，帮助完成数据初始化
	IFilterConfigSlice = []entity.Filter{
		{
			Model:   gorm.Model{ID: 2},
			Name:    Limiter,
			Remarks: "限流过滤器",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    NameList,
			Remarks: "黑白名单过滤器",
		},
	}

	// 注册到Map中，帮助完成网关流程
	IFilterByNameMap[Limiter] = LimiterFilter
	IFilterByNameMap[NameList] = NameListFilter
}
