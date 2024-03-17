package final

import (
	"encoding/json"
	"fmt"
	"loiter/app/store"
	"loiter/global"
	"loiter/model/entity"
	"loiter/utils"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

/**
 * 请求日志插件
 * @auth eyesYeager
 * @date 2024/2/13 16:16
 */

var RequestLogFinal = requestLogFinal{
	bufferMaxLen: 100,
}

type requestLogFinal struct {
	bufferMaxLen int
}

var logBufferList []entity.RequestLog

// RequestLogFinal 请求日志插件
func (r *requestLogFinal) RequestLogFinal(_ http.ResponseWriter, req *http.Request, _ *http.Response, host string, entrance string, errInfo string) error {
	endTime := time.Now().UnixMilli()
	endTimeStr := strconv.FormatInt(endTime, 10)
	requestInfo, _ := httputil.DumpRequest(req, true)
	startTimeStr := store.GetValue(req, store.RequestBeginTime)
	startTime, _ := strconv.ParseInt(startTimeStr, 10, 64)
	requestLogEntity := entity.RequestLog{
		RequestId: store.GetValue(req, store.RequestId),
		Host:      host,
		Path:      req.URL.Path,
		ReqInfo:   string(requestInfo),
		Ip:        utils.GetIp(req),
		Browser:   utils.GetBrowser(req),
		StartTime: startTimeStr,
		EndTime:   endTimeStr,
		RunTime:   uint64(endTime - startTime),
		Entrance:  entrance,
		ErrInfo:   errInfo,
	}
	logBufferList = append(logBufferList, requestLogEntity)
	if len(logBufferList) >= r.bufferMaxLen {
		//
	}
	go func() {
		if err := global.MDB.Create(&requestLogEntity).Error; err != nil {
			requestLogEntityJson, _ := json.Marshal(requestLogEntity)
			global.GatewayLogger.Error(fmt.Sprintf("RequestLogFinal plugin execution failed! RequestLogEntity: %s, error: %s", string(requestLogEntityJson), err.Error()))
		}
	}()
	return nil
}
