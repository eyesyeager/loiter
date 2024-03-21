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
	"sync"
	"time"
)

/**
 * 请求日志插件
 * @auth eyesYeager
 * @date 2024/2/13 16:16
 */

func init() {
	RequestLogFinal.startTick()
}

var RequestLogFinal = requestLogFinal{
	ticker:        time.NewTicker(time.Duration(10) * time.Second),
	bufferMaxLen:  100,
	logBufferList: []entity.RequestLog{},
}

type requestLogFinal struct {
	mutex         sync.Mutex
	ticker        *time.Ticker
	bufferMaxLen  int
	logBufferList []entity.RequestLog
}

// RequestLogFinal 请求日志插件
func (r *requestLogFinal) RequestLogFinal(_ http.ResponseWriter, req *http.Request, _ *http.Response, host string, entrance string, errInfo string) error {
	requestInfo, _ := httputil.DumpRequest(req, true)
	startTimeStr := store.GetValue(req, store.RequestBeginTime)
	startTime, _ := strconv.ParseInt(startTimeStr, 10, 64)
	endTime := time.Now().UnixMilli()
	endTimeStr := strconv.FormatInt(endTime, 10)
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
	r.logBufferList = append(r.logBufferList, requestLogEntity)
	if len(r.logBufferList) >= r.bufferMaxLen {
		r.persistence()
	}
	return nil
}

// startTick 开启定时任务
func (r *requestLogFinal) startTick() {
	go func() {
		for range r.ticker.C {
			r.persistence()
		}
	}()
}

// persistence 持久化
func (r *requestLogFinal) persistence() {
	if len(r.logBufferList) == 0 {
		return
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	container := r.logBufferList
	r.logBufferList = []entity.RequestLog{}
	go func() {
		if err := global.MDB.Create(&container).Error; err != nil {
			requestLogEntityJson, _ := json.Marshal(container)
			global.GatewayLogger.Error(fmt.Sprintf("RequestLogFinal plugin execution failed! RequestLogEntity: %s, error: %s", string(requestLogEntityJson), err.Error()))
		}
	}()
}
