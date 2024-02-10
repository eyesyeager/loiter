package proxy

import (
	"fmt"
	"github.com/rs/xid"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/filter"
	"loiter/plugin/store"
	"net/http"
	"strconv"
	"time"
)

/**
 * 代理前置处理
 * @auth eyesYeager
 * @date 2024/1/29 15:32
 */

// pre 前置处理总入口
func pre(w http.ResponseWriter, r *http.Request, host string) bool {
	buildRequestId(r)
	if allow := entryFilter(w, r, host); !allow {
		return allow
	}
	return true
}

// buildRequestId 生成全局唯一请求id
func buildRequestId(r *http.Request) {
	// 写入日志Id，用于填充请求日志
	if err := store.SetValue(r, store.RequestId, xid.New().String()); err != nil {
		panic(fmt.Errorf("write request id failed, error: %s", err.Error()))
	}
	// 写入请求时间，用于计算请求耗时
	if err := store.SetValue(r, store.RequestBeginTime, strconv.FormatInt(time.Now().UnixMilli(), 10)); err != nil {
		panic(fmt.Errorf("write request begin time failed, error: %s", err.Error()))
	}
}

// entryFilter 进入过滤器
func entryFilter(w http.ResponseWriter, r *http.Request, host string) bool {
	// 进入过滤器
	err, allow := filter.Entry(w, r, host)
	if err != nil {
		errMsg := fmt.Sprintf("filter execution failed. Error message: %s", err.Error())
		statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.BadGateway, errMsg)
		helper.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(errMsg)
		return false
	}
	return allow
}
