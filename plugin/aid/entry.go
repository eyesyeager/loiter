package aid

import (
	"gorm.io/gorm"
	"loiter/kernel/model/entity"
	"net/http"
)

/**
 * 响应处理器入口
 * @auth eyesYeager
 * @date 2024/1/25 17:17
 */

const (
	RequestFill = "requestFill"
)

// IAid 响应处理器接口
type IAid func(http.ResponseWriter, *http.Request, *http.Response, string) error

// IAidByNameMap 请求通道方法列表 by 通道名
var IAidByNameMap = make(map[string]IAid)

// IAidConfigSlice 请求通道策略切片
var IAidConfigSlice []entity.Aid

// Register 注册插件
func Register() {
	// 注册到Map中，帮助完成网关流程
	IAidByNameMap[RequestFill] = RequestFillAid

	// 注册到config中，帮助完成数据初始化
	IAidConfigSlice = []entity.Aid{
		{
			Model:   gorm.Model{ID: 1},
			Name:    RequestFill,
			Remarks: "填充请求日志",
		},
	}
}
