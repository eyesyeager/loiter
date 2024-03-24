package container

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/config"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"strings"
)

/**
 * 处理器容器
 * @auth eyesYeager
 * @date 2024/2/12 19:09
 */

// FilterByAppMap 过滤器有序列表 by AppHost
var FilterByAppMap = make(map[string][]string)

// AidByAppMap 响应处理器 by AppHost
var AidByAppMap = make(map[string][]string)

// ExceptionByAppMap 异常处理器 by AppHost
var ExceptionByAppMap = make(map[string][]string)

// FinalByAppMap 最终处理器 by AppHost
var FinalByAppMap = make(map[string][]string)

// InitProcessor 初始化处理器容器
func InitProcessor() {
	global.AppLogger.Info("start initializing the Processor container")
	// 获取处理器配置
	var appProcessorNameList []po.GetAppProcessorName
	if rowsAffected := global.MDB.Raw(`SELECT a.host, ap.genre, ap.codes
						FROM app a, app_processor ap
						WHERE a.status = ? AND a.id = ap.app_id`, constant.Status.Normal.Code).Scan(&appProcessorNameList).RowsAffected; rowsAffected == 0 {
		global.AppLogger.Info("there is currently no valid Processor configuration")
		return
	}
	// 构建处理器临时容器
	var filterByAppTempMap = make(map[string][]string)
	var aidByAppTempMap = make(map[string][]string)
	var exceptionByAppTempMap = make(map[string][]string)
	var finalByAppTempMap = make(map[string][]string)
	for _, item := range appProcessorNameList {
		if item.Codes == "" {
			continue
		}
		processorList := strings.Split(item.Codes, config.Program.PluginConfig.ProcessorDelimiter)
		if item.Genre == constants.Processor.Filter.Code {
			filterByAppTempMap[item.Host] = processorList
		} else if item.Genre == constants.Processor.Aid.Code {
			aidByAppTempMap[item.Host] = processorList
		} else if item.Genre == constants.Processor.Exception.Code {
			exceptionByAppTempMap[item.Host] = processorList
		} else if item.Genre == constants.Processor.Final.Code {
			finalByAppTempMap[item.Host] = processorList
		} else {
			global.AppLogger.Error(fmt.Sprintf("when initializing the container, it was found that the application with host %s has an illegal processor type: %s", item.Host, item.Genre))
		}
	}
	FilterByAppMap = filterByAppTempMap
	AidByAppMap = aidByAppTempMap
	ExceptionByAppMap = exceptionByAppTempMap
	FinalByAppMap = finalByAppTempMap
	global.AppLogger.Info("complete the initialization of Processor container")
}

// RefreshProcessor 刷新处理器容器
func RefreshProcessor(appId uint) error {
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		return errors.New(fmt.Sprintf("appId为%d的应用刷新处理器容器失败，应用不存在！", appId))
	}
	// 如果状态不为normal，则清空容器项
	if checkApp.Status != constant.Status.Normal.Code {
		DeleteProcessor(checkApp.Host)
		return nil
	}
	var appProcessorNameList []po.GetAppProcessorName
	if err := global.MDB.Raw(`SELECT a.host, ap.genre, ap.codes
						FROM app a, app_processor ap
						WHERE a.id = ap.app_id AND a.id = ?`, appId).Scan(&appProcessorNameList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 如果不存在处理器，则清空容器项
	if len(appProcessorNameList) == 0 {
		DeleteProcessor(checkApp.Host)
		return nil
	}
	// 如果存在处理器，则尝试刷新容器项
	processorMap := make(map[string]string)
	for _, item := range appProcessorNameList {
		processorMap[item.Genre] = item.Codes
	}
	refreshProcessorChild(processorMap, constants.Processor.Filter.Code, checkApp.Host)
	refreshProcessorChild(processorMap, constants.Processor.Aid.Code, checkApp.Host)
	refreshProcessorChild(processorMap, constants.Processor.Exception.Code, checkApp.Host)
	refreshProcessorChild(processorMap, constants.Processor.Final.Code, checkApp.Host)
	return nil
}

// DeleteProcessor 删除处理器容器
func DeleteProcessor(host string) {
	delete(FilterByAppMap, host)
	delete(AidByAppMap, host)
	delete(ExceptionByAppMap, host)
	delete(FinalByAppMap, host)
}

// refreshProcessorChild 刷新处理器子项
func refreshProcessorChild(processorMap map[string]string, genre string, host string) {
	var m map[string][]string
	if genre == constants.Processor.Filter.Code {
		m = FilterByAppMap
	} else if genre == constants.Processor.Aid.Code {
		m = AidByAppMap
	} else if genre == constants.Processor.Exception.Code {
		m = ExceptionByAppMap
	} else {
		m = FinalByAppMap
	}
	item, ok := processorMap[genre]
	if ok {
		m[host] = strings.Split(item, config.Program.PluginConfig.ProcessorDelimiter)
	} else {
		delete(m, host)
	}
}
