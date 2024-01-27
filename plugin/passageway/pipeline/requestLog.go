package pipeline

import (
	"encoding/json"
	"fmt"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/model/entity"
	"loiter/kernel/store"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

/**
 * 记录请求日志
 * @auth eyesYeager
 * @date 2024/1/11 16:45
 */

func RequestLogPipeline(_ http.ResponseWriter, r *http.Request, host string) (error, bool) {
	// 写入请求时间，用于计算请求耗时
	if err := store.SetValue(r, store.RequestLogTime, strconv.FormatInt(time.Now().UnixMilli(), 10)); err != nil {
		global.GatewayLogger.Error(fmt.Sprintf("write request time failed, error: %s", err.Error()))
	}
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
		// 写入日志Id，用于填充请求日志
		if err := store.SetValue(r, store.RequestLogId, strconv.Itoa(int(requestLogEntity.ID))); err != nil {
			global.GatewayLogger.Error(fmt.Sprintf("write request log id failed, error: %s", err.Error()))
		}
	}()
	return nil, true
}
