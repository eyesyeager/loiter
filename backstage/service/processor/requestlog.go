package processor

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"strconv"
	"time"
)

/**
 * 请求日志业务层
 * @auth eyesYeager
 * @date 2024/2/29 17:54
 */

type requestLogService struct {
}

var RequestLogService = requestLogService{}

// GetOverviewRequestLog 获取请求日志概览
func (*requestLogService) GetOverviewRequestLog() (error, returnee.GetOverviewRequestLog) {
	nowDayTime := time.Now().Format(time.DateOnly)
	res := returnee.GetOverviewRequestLog{}
	if err := global.MDB.Raw(`SELECT COUNT(*) RequestNum, FLOOR(AVG(run_time)) AvgRunTime, COUNT(DISTINCT ip) VisitorNum 
				FROM request_log WHERE created_at > ?`, nowDayTime).Scan(&res).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	var rejectNum uint64
	if err := global.MDB.Raw(`SELECT COUNT(*) FROM request_log WHERE created_at > ? AND entrance = ?`, nowDayTime, "reject").Scan(&rejectNum).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	res.RejectNum = rejectNum
	return nil, res
}

// GetDetailedRequestExtremumLog 获取请求日志详细信息-极值数据
func (*requestLogService) GetDetailedRequestExtremumLog(data receiver.GetDetailedRequestLog) (err error, res returnee.GetDetailedRequestExtremumLog) {
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err = global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), res
		}
	}
	// 时间边界处理
	var dateFrom, dateTo string
	dayNow := time.Now().Format(time.DateOnly)
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dateFrom = dayNow
		dateTo = dayNow
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		dateFrom = dayYesterday
		dateTo = dayYesterday
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dateFrom = time.Now().AddDate(0, 0, -6).Format(time.DateOnly)
		dateTo = dayNow
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dateFrom = time.Now().AddDate(0, 0, -29).Format(time.DateOnly)
		dateTo = dayNow
	} else {
		return errors.New(fmt.Sprintf("非法时间间隔字段: %s", data.TimeInterval)), res
	}
	dateFrom += " 00:00:00"
	dateTo += " 23:59:59"
	// 获取请求数、响应时间相关数据
	if err = global.MDB.Raw(`SELECT COUNT(*) RequestNum, MIN(run_time) RunTimeMin, MAX(run_time) RunTimeMax, FLOOR(AVG(run_time)) RunTimeAvg
			FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?)`, dateFrom, dateTo, checkApp.Host, checkApp.Host).Scan(&res).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 计算平均 QPS
	loc, _ := time.LoadLocation("Local")
	var dayFromDate time.Time
	var dayToDate time.Time
	dayFromDate, _ = time.ParseInLocation(time.DateTime, dateFrom, loc)
	dayToDate, _ = time.ParseInLocation(time.DateTime, dateTo, loc)
	res.QPSAvg = res.RequestNum / (dayToDate.Unix() - dayFromDate.Unix())
	// 获取请求拒绝数
	var rejectNum int64
	if err = global.MDB.Raw(`SELECT COUNT(*) FROM request_log WHERE created_at > ? AND created_at < ? AND entrance IN ? AND ('' = ? OR host = ?)`,
		dateFrom, dateTo, []string{constants.PostEntrance.Reject, constants.PostEntrance.Error}, checkApp.Host, checkApp.Host).Scan(&rejectNum).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	res.RequestReject = rejectNum
	return err, res
}

// GetDetailedRequestNumLog 获取请求日志详细信息-请求数据
func (*requestLogService) GetDetailedRequestNumLog(data receiver.GetDetailedRequestLog) (error, returnee.GetDetailedRequestNumLog) {
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), returnee.GetDetailedRequestNumLog{}
		}
	}
	// 构建请求数据
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dayNow := time.Now().Format(time.DateOnly)
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestNumDayLog(checkApp.Host, dayNow, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		err, _, lastSeries := getDetailedRequestNumDayLog(checkApp.Host, dayYesterday, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		return nil, returnee.GetDetailedRequestNumLog{XAxis: xAxis, Series: series, LastSeries: lastSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		dayLast := time.Now().AddDate(0, 0, -2).Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestNumDayLog(checkApp.Host, dayYesterday, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		err, _, lastSeries := getDetailedRequestNumDayLog(checkApp.Host, dayLast, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		return nil, returnee.GetDetailedRequestNumLog{XAxis: xAxis, Series: series, LastSeries: lastSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestNumWeekLog(checkApp.Host, dayNow, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		return nil, returnee.GetDetailedRequestNumLog{XAxis: xAxis, Series: series}
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestNumMonthLog(checkApp.Host, dayNow, "")
		if err != nil {
			return err, returnee.GetDetailedRequestNumLog{}
		}
		return nil, returnee.GetDetailedRequestNumLog{XAxis: xAxis, Series: series}
	} else {
		return errors.New(fmt.Sprintf("非法时间范围字段: %s", data.TimeInterval)), returnee.GetDetailedRequestNumLog{}
	}
}

// GetDetailedRequestRuntimeLog 获取请求日志详细信息-响应时间
func (*requestLogService) GetDetailedRequestRuntimeLog(data receiver.GetDetailedRequestLog) (error, returnee.GetDetailedRequestRuntimeLog) {
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), returnee.GetDetailedRequestRuntimeLog{}
		}
	}
	// 构建请求数据
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestRuntimeDayLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestRuntimeLog{}
		}
		return nil, returnee.GetDetailedRequestRuntimeLog{XAxis: xAxis, Series: series}
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestRuntimeDayLog(checkApp.Host, dayYesterday)
		if err != nil {
			return err, returnee.GetDetailedRequestRuntimeLog{}
		}
		return nil, returnee.GetDetailedRequestRuntimeLog{XAxis: xAxis, Series: series}
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestRuntimeWeekLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestRuntimeLog{}
		}
		return nil, returnee.GetDetailedRequestRuntimeLog{XAxis: xAxis, Series: series}
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestRuntimeMonthLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestRuntimeLog{}
		}
		return nil, returnee.GetDetailedRequestRuntimeLog{XAxis: xAxis, Series: series}
	} else {
		return errors.New(fmt.Sprintf("非法时间范围字段: %s", data.TimeInterval)), returnee.GetDetailedRequestRuntimeLog{}
	}
}

// getDetailedRequestRuntimeDayLog 获取请求日志详细信息-响应时间-日
func getDetailedRequestRuntimeDayLog(host string, day string) (error, []string, []int) {
	xAxis := getDayXAxis()
	// 请求总数
	var detailedRuntimeList []po.GetDetailedRuntime
	if err := global.MDB.Raw(`SELECT FLOOR(AVG(run_time)) Runtime, DATE_FORMAT(created_at, '%H') Time 
						FROM request_log WHERE DATE_FORMAT(created_at, '%Y-%m-%d') = ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, day, host, host).Scan(&detailedRuntimeList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var listCur = 0
	var series []int
	for i := 0; i < 24; i++ {
		var hour string
		if i < 10 {
			hour = "0" + strconv.Itoa(i)
		} else {
			hour = strconv.Itoa(i)
		}
		if detailedRuntimeList == nil || listCur >= len(detailedRuntimeList) || detailedRuntimeList[listCur].Time != hour {
			series = append(series, -1)
		} else {
			series = append(series, detailedRuntimeList[listCur].Runtime)
			listCur++
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestRuntimeWeekLog 获取请求日志详细信息-响应时间-周
func getDetailedRequestRuntimeWeekLog(host string, day string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getWeekXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -6).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedRuntimeList []po.GetDetailedRuntime
	if err := global.MDB.Raw(`SELECT FLOOR(AVG(run_time)) Runtime, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host).Scan(&detailedRuntimeList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedRuntimeList == nil || listCur >= len(detailedRuntimeList) {
			series = append(series, -1)
			continue
		}
		noYearDate := detailedRuntimeList[listCur].Time[5:]
		if noYearDate == item {
			series = append(series, detailedRuntimeList[listCur].Runtime)
			listCur++
		} else {
			series = append(series, -1)
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestRuntimeMonthLog 获取请求日志详细信息-响应时间-月
func getDetailedRequestRuntimeMonthLog(host string, day string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getFullMonthXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -29).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedRuntimeList []po.GetDetailedRuntime
	if err := global.MDB.Raw(`SELECT FLOOR(AVG(run_time)) Runtime, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host).Scan(&detailedRuntimeList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedRuntimeList == nil || listCur >= len(detailedRuntimeList) {
			series = append(series, -1)
			continue
		}
		if item == detailedRuntimeList[listCur].Time {
			series = append(series, detailedRuntimeList[listCur].Runtime)
			listCur++
		} else {
			series = append(series, -1)
		}
	}
	// 横坐标处理
	dealMonthXAxis(xAxis)
	return nil, xAxis, series
}

// GetDetailedRequestQPSLog 获取请求日志详细信息-QPS
func (p *requestLogService) GetDetailedRequestQPSLog(data receiver.GetDetailedRequestLog) (error, returnee.GetDetailedRequestQPSLog) {
	err, numRes := p.GetDetailedRequestNumLog(data)
	if err != nil {
		return err, returnee.GetDetailedRequestQPSLog{}
	}
	if data.TimeInterval == constant.RequestLogInterval.Today || data.TimeInterval == constant.RequestLogInterval.Yesterday {
		for index := range numRes.XAxis {
			numRes.Series[index] /= constant.Time.HourSeconds
			numRes.LastSeries[index] /= constant.Time.HourSeconds
		}
	} else {
		for index := range numRes.XAxis {
			numRes.Series[index] /= constant.Time.DaySeconds
		}
	}
	return nil, returnee.GetDetailedRequestQPSLog{XAxis: numRes.XAxis, Series: numRes.Series, LastSeries: numRes.LastSeries}
}

// GetDetailedRequestVisitorLog 获取请求日志详细信息-访客数据
func (*requestLogService) GetDetailedRequestVisitorLog(data receiver.GetDetailedRequestLog) (error, returnee.GetDetailedRequestVisitorLog) {
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), returnee.GetDetailedRequestVisitorLog{}
		}
	}
	// 构建请求数据
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dayNow := time.Now().Format(time.DateOnly)
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestVisitorDayLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		err, _, lastSeries := getDetailedRequestVisitorDayLog(checkApp.Host, dayYesterday)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		return nil, returnee.GetDetailedRequestVisitorLog{XAxis: xAxis, Series: series, LastSeries: lastSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		dayLast := time.Now().AddDate(0, 0, -2).Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestVisitorDayLog(checkApp.Host, dayYesterday)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		err, _, lastSeries := getDetailedRequestVisitorDayLog(checkApp.Host, dayLast)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		return nil, returnee.GetDetailedRequestVisitorLog{XAxis: xAxis, Series: series, LastSeries: lastSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestVisitorWeekLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		return nil, returnee.GetDetailedRequestVisitorLog{XAxis: xAxis, Series: series}
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, series := getDetailedRequestVisitorMonthLog(checkApp.Host, dayNow)
		if err != nil {
			return err, returnee.GetDetailedRequestVisitorLog{}
		}
		return nil, returnee.GetDetailedRequestVisitorLog{XAxis: xAxis, Series: series}
	} else {
		return errors.New(fmt.Sprintf("非法时间范围字段: %s", data.TimeInterval)), returnee.GetDetailedRequestVisitorLog{}
	}
}

// getDetailedRequestVisitorDayLog 获取请求日志详细信息-访客数据-日
func getDetailedRequestVisitorDayLog(host string, day string) (error, []string, []int) {
	xAxis := getDayXAxis()
	// 请求总数
	var detailedDayRequestList []po.GetDetailedDayRequest
	if err := global.MDB.Raw(`SELECT COUNT(DISTINCT ip) Num, DATE_FORMAT(created_at, '%H') Time 
						FROM request_log WHERE DATE_FORMAT(created_at, '%Y-%m-%d') = ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, day, host, host).Scan(&detailedDayRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var listCur = 0
	var series []int
	for i := 0; i < 24; i++ {
		var hour string
		if i < 10 {
			hour = "0" + strconv.Itoa(i)
		} else {
			hour = strconv.Itoa(i)
		}
		if detailedDayRequestList == nil || listCur >= len(detailedDayRequestList) || detailedDayRequestList[listCur].Time != hour {
			series = append(series, 0)
		} else {
			series = append(series, detailedDayRequestList[listCur].Num)
			listCur++
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestVisitorWeekLog 获取请求日志详细信息-访客数据-周
func getDetailedRequestVisitorWeekLog(host string, day string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getWeekXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -6).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedWeekRequestList []po.GetDetailedWeekRequest
	if err := global.MDB.Raw(`SELECT COUNT(DISTINCT ip) Num, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host).Scan(&detailedWeekRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedWeekRequestList == nil || listCur >= len(detailedWeekRequestList) {
			series = append(series, 0)
			continue
		}
		noYearDate := detailedWeekRequestList[listCur].Time[5:]
		if noYearDate == item {
			series = append(series, detailedWeekRequestList[listCur].Num)
			listCur++
		} else {
			series = append(series, 0)
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestVisitorMonthLog 获取请求日志详细信息-访客数据-月
func getDetailedRequestVisitorMonthLog(host string, day string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getFullMonthXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -29).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedMonthRequestList []po.GetDetailedMonthRequest
	if err := global.MDB.Raw(`SELECT COUNT(DISTINCT ip) Num, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) 
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host).Scan(&detailedMonthRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedMonthRequestList == nil || listCur >= len(detailedMonthRequestList) {
			series = append(series, 0)
			continue
		}
		if item == detailedMonthRequestList[listCur].Time {
			series = append(series, detailedMonthRequestList[listCur].Num)
			listCur++
		} else {
			series = append(series, 0)
		}
	}
	// 横坐标处理
	dealMonthXAxis(xAxis)
	return nil, xAxis, series
}

// GetDetailedRequestTopApiLog 获取请求日志详情-Top接口
func (*requestLogService) GetDetailedRequestTopApiLog(data receiver.GetDetailedRequestLog) (error, []returnee.GetDetailedRequestTopApiLog) {
	limitNum := 10
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), nil
		}
	}
	// 时间处理
	var dateFrom, dateTo string
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dayNow := time.Now().Format(time.DateOnly)
		dateFrom = dayNow
		dateTo = dayNow
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		dateFrom = dayYesterday
		dateTo = dayYesterday
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dateFrom = time.Now().AddDate(0, 0, -6).Format(time.DateOnly)
		dateTo = time.Now().Format(time.DateOnly)
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dateFrom = time.Now().AddDate(0, 0, -29).Format(time.DateOnly)
		dateTo = time.Now().Format(time.DateOnly)
	} else {
		return errors.New(fmt.Sprintf("非法时间范围字段: %s", data.TimeInterval)), nil
	}
	dateFrom += " 00:00:00"
	dateTo += " 23:59:59"
	// 获取数据
	var res []returnee.GetDetailedRequestTopApiLog
	if err := global.MDB.Raw(`SELECT CONCAT(host, path) Api, COUNT(*) Num FROM request_log 
                        	WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?)
                            GROUP BY host, path ORDER BY Num DESC LIMIT ?`, dateFrom, dateTo, checkApp.Host, checkApp.Host, limitNum).Scan(&res).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	if res == nil || len(res) == 0 {
		return nil, res
	}
	var totalNum int64
	if err := global.MDB.Raw(`SELECT COUNT(*) FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?)`,
		dateFrom, dateTo, checkApp.Host, checkApp.Host).Scan(&totalNum).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for index := range res {
		res[index].Rate = fmt.Sprintf("%.2f", float64(res[index].Num)/float64(totalNum)*100) + "%"
	}
	return nil, res
}

// GetDetailedRequestRejectLog 获取请求日志详情-请求拒绝数
func (*requestLogService) GetDetailedRequestRejectLog(data receiver.GetDetailedRequestLog) (error, returnee.GetDetailedRequestRejectLog) {
	// 获取 host 信息
	var checkApp = entity.App{Name: data.AppName}
	if data.AppName != "" {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("网络异常或应用 %s 不存在", data.AppName)), returnee.GetDetailedRequestRejectLog{}
		}
	}
	// 构建请求数据
	if data.TimeInterval == constant.RequestLogInterval.Today { // 计算今天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, rejectSeries := getDetailedRequestNumDayLog(checkApp.Host, dayNow, constants.PostEntrance.Reject)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		err, _, errorSeries := getDetailedRequestNumDayLog(checkApp.Host, dayNow, constants.PostEntrance.Error)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		return nil, returnee.GetDetailedRequestRejectLog{XAxis: xAxis, RejectSeries: rejectSeries, ErrorSeries: errorSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Yesterday { // 计算昨天数据
		dayYesterday := time.Now().AddDate(0, 0, -1).Format(time.DateOnly)
		err, xAxis, rejectSeries := getDetailedRequestNumDayLog(checkApp.Host, dayYesterday, constants.PostEntrance.Reject)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		err, _, errorSeries := getDetailedRequestNumDayLog(checkApp.Host, dayYesterday, constants.PostEntrance.Error)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		return nil, returnee.GetDetailedRequestRejectLog{XAxis: xAxis, RejectSeries: rejectSeries, ErrorSeries: errorSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Week { // 计算近7天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, rejectSeries := getDetailedRequestNumWeekLog(checkApp.Host, dayNow, constants.PostEntrance.Reject)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		err, _, errorSeries := getDetailedRequestNumWeekLog(checkApp.Host, dayNow, constants.PostEntrance.Error)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		return nil, returnee.GetDetailedRequestRejectLog{XAxis: xAxis, RejectSeries: rejectSeries, ErrorSeries: errorSeries}
	} else if data.TimeInterval == constant.RequestLogInterval.Month { // 计算近30天数据
		dayNow := time.Now().Format(time.DateOnly)
		err, xAxis, rejectSeries := getDetailedRequestNumMonthLog(checkApp.Host, dayNow, constants.PostEntrance.Reject)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		err, _, errorSeries := getDetailedRequestNumMonthLog(checkApp.Host, dayNow, constants.PostEntrance.Error)
		if err != nil {
			return err, returnee.GetDetailedRequestRejectLog{}
		}
		return nil, returnee.GetDetailedRequestRejectLog{XAxis: xAxis, RejectSeries: rejectSeries, ErrorSeries: errorSeries}
	} else {
		return errors.New(fmt.Sprintf("非法时间范围字段: %s", data.TimeInterval)), returnee.GetDetailedRequestRejectLog{}
	}
}

/***********************************************************************
 *                              help
 ***********************************************************************/

// getDetailedRequestNumDayLog 获取请求日志详细信息-请求数据-日
func getDetailedRequestNumDayLog(host string, day string, entrance string) (error, []string, []int) {
	xAxis := getDayXAxis()
	// 请求总数
	var detailedDayRequestList []po.GetDetailedDayRequest
	if err := global.MDB.Raw(`SELECT COUNT(*) Num, DATE_FORMAT(created_at, '%H') Time 
						FROM request_log WHERE DATE_FORMAT(created_at, '%Y-%m-%d') = ? AND ('' = ? OR host = ?) AND ('' = ? OR entrance = ?)
		                GROUP BY Time ORDER BY Time`, day, host, host, entrance, entrance).Scan(&detailedDayRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var listCur = 0
	var series []int
	for i := 0; i < 24; i++ {
		var hour string
		if i < 10 {
			hour = "0" + strconv.Itoa(i)
		} else {
			hour = strconv.Itoa(i)
		}
		if detailedDayRequestList == nil || listCur >= len(detailedDayRequestList) || detailedDayRequestList[listCur].Time != hour {
			series = append(series, 0)
		} else {
			series = append(series, detailedDayRequestList[listCur].Num)
			listCur++
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestNumWeekLog 获取请求日志详细信息-请求数据-周
func getDetailedRequestNumWeekLog(host string, day string, entrance string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getWeekXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -6).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedWeekRequestList []po.GetDetailedWeekRequest
	if err := global.MDB.Raw(`SELECT COUNT(*) Num, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) AND ('' = ? OR entrance = ?)
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host, entrance, entrance).Scan(&detailedWeekRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedWeekRequestList == nil || listCur >= len(detailedWeekRequestList) {
			series = append(series, 0)
			continue
		}
		noYearDate := detailedWeekRequestList[listCur].Time[5:]
		if noYearDate == item {
			series = append(series, detailedWeekRequestList[listCur].Num)
			listCur++
		} else {
			series = append(series, 0)
		}
	}
	return nil, xAxis, series
}

// getDetailedRequestNumMonthLog 获取请求日志详细信息-请求数据-月
func getDetailedRequestNumMonthLog(host string, day string, entrance string) (error, []string, []int) {
	// 横坐标组装
	dayFront, _ := time.Parse(time.DateOnly, day)
	xAxis := getFullMonthXAxis(dayFront)
	// 数据组装
	dateFrom := dayFront.AddDate(0, 0, -29).Format(time.DateOnly) + " 00:00:00"
	dateTo := day + " 23:59:59"
	var detailedMonthRequestList []po.GetDetailedMonthRequest
	if err := global.MDB.Raw(`SELECT COUNT(*) Num, DATE_FORMAT(created_at, '%Y-%m-%d') Time
						FROM request_log WHERE created_at >= ? AND created_at <= ? AND ('' = ? OR host = ?) AND ('' = ? OR entrance = ?)
		                GROUP BY Time ORDER BY Time`, dateFrom, dateTo, host, host, entrance, entrance).Scan(&detailedMonthRequestList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil, nil
	}
	var series []int
	var listCur = 0
	for _, item := range xAxis {
		if detailedMonthRequestList == nil || listCur >= len(detailedMonthRequestList) {
			series = append(series, 0)
			continue
		}
		if item == detailedMonthRequestList[listCur].Time {
			series = append(series, detailedMonthRequestList[listCur].Num)
			listCur++
		} else {
			series = append(series, 0)
		}
	}
	// 横坐标处理
	dealMonthXAxis(xAxis)
	return nil, xAxis, series
}

// getDayXAxis 获取统计表格横坐标-日
func getDayXAxis() []string {
	return []string{"0", "", "", "", "4", "", "", "", "8", "", "", "", "12", "", "", "", "16", "", "", "", "20", "", "", ""}
}

// getWeekXAxis 获取统计表格横坐标-周
func getWeekXAxis(day time.Time) []string {
	var xAxis []string
	for i := 6; i >= 0; i-- {
		xAxis = append(xAxis, day.AddDate(0, 0, -i).Format("01-02"))
	}
	return xAxis
}

// getFullMonthXAxis 获取统计表格横坐标-月（完整数据）
func getFullMonthXAxis(day time.Time) []string {
	var xAxis []string
	for i := 0; i < 30; i++ {
		xAxis = append(xAxis, day.AddDate(0, 0, -29+i).Format(time.DateOnly))
	}
	return xAxis
}

// dealMonthXAxis 处理统计表格横坐标-月
func dealMonthXAxis(xAxis []string) {
	for i := 0; i < 30; i++ {
		if i%5 != 0 {
			xAxis[i] = ""
		}
	}
}
