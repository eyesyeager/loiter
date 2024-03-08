package exception

import (
	"loiter/constants"
	"loiter/model/entity"
	"net/http"
)

/**
 * 异常插件
 * @auth eyesYeager
 * @date 2024/1/30 09:59
 */

const (
	Notice = "notice"
)

// IException 异常处理器接口
type IException func(http.ResponseWriter, *http.Request, string, string)

// IExceptionByNameMap 异常处理器方法列表 by 处理器名
var IExceptionByNameMap = make(map[string]IException)

// IExceptionConfigList 响应处理器切片
var IExceptionConfigList []entity.Processor

// Register 注册插件
func Register() {
	IExceptionConfigList = []entity.Processor{
		{
			Code:    Notice,
			Name:    "异常通知",
			Genre:   constants.Processor.Exception.Code,
			Remarks: "若过滤器出现异常，可由此插件发送站内消息与邮件消息。",
		},
	}

	IExceptionByNameMap[Notice] = NoticeException
}
