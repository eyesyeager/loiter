package final

import (
	"loiter/constants"
	"loiter/model/entity"
	"net/http"
)

/**
 * 最终处理器插件
 * @auth eyesYeager
 * @date 2024/1/30 09:56
 */

const (
	RequestLog = "requestLog"
)

// IFinal 最终处理器接口
type IFinal func(http.ResponseWriter, *http.Request, *http.Response, string, string, string) error

// IFinalByNameMap 最终处理器方法列表 by 处理器名
var IFinalByNameMap = make(map[string]IFinal)

// IFinalConfigList 最终处理器切片
var IFinalConfigList []entity.Processor

// Register 注册插件
func Register() {
	IFinalConfigList = []entity.Processor{
		{
			Code:    RequestLog,
			Name:    "请求日志",
			Genre:   constants.Processor.Final.Code,
			Remarks: "本插件是将请求信息持久化到数据库中，因此不适用于大流量应用。管理系统中大部分数据统计功能都是基于本插件。",
		},
	}

	IFinalByNameMap[RequestLog] = RequestLogFinal
}
