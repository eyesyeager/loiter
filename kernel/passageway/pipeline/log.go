package pipeline

import (
	"encoding/json"
	"loiter/global"
	"loiter/kernel/helper"
	"loiter/kernel/model/entity"
	"net/http"
	"net/http/httputil"
)

/**
 * 日志管道
 * @auth eyesYeager
 * @date 2024/1/11 16:45
 */

func RequestLogPipeline(_ http.ResponseWriter, r *http.Request, host string) (error, bool) {
	go func() {
		requestInfo, _ := httputil.DumpRequest(r, true)
		requestLogEntity := entity.RequestLog{
			Host:    host,
			Path:    r.URL.Path,
			ReqInfo: string(requestInfo),
			Ip:      helper.GetIp(r),
			Browser: helper.GetBrowser(r),
		}
		if err := global.MDB.Create(&requestLogEntity).Error; err != nil {
			requestLogEntityJson, _ := json.Marshal(requestLogEntity)
			global.BackstageLogger.Error("passageway-pipeline-RequestLog-Create error!",
				"requestLogEntity:", string(requestLogEntityJson),
				";error:", err.Error())
		}
	}()
	return nil, true
}
