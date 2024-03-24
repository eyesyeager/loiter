package service

import (
	"errors"
	"fmt"
	"loiter/app/plugin/filter/limiter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/foundation"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/returnee"
	"reflect"
	"strconv"
)

/**
 * 通用数据业务
 * @auth eyesYeager
 * @date 2024/2/22 11:56
 */

type commonService struct {
}

var CommonService = commonService{}

// GetStatusDictionary 获取状态字典
func (*commonService) GetStatusDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constant.Status)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		statusStructure := val.(constant.StatusStructure)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: statusStructure.Name,
			Value: strconv.Itoa(int(statusStructure.Code)),
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetAppDictionary 获取应用字典
func (*commonService) GetAppDictionary() (error, []returnee.GetDictionary) {
	var appList []entity.App
	if err := global.MDB.Find(&appList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil
	}
	var dictionaryList []returnee.GetDictionary
	for _, item := range appList {
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: item.Name,
			Value: strconv.Itoa(int(item.ID)),
		})
	}
	// 如果没有应用，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetBalancerDictionary 获取负载均衡字典
func (*commonService) GetBalancerDictionary() (error, []returnee.GetDictionary) {
	var balancerList []entity.Balancer
	if err := global.MDB.Find(&balancerList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil
	}
	var dictionaryList []returnee.GetDictionary
	for _, item := range balancerList {
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: item.Name,
			Value: item.Code,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetNoticeDictionary 获取通知类型字典
func (*commonService) GetNoticeDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constants.Notice)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		value := val.(string)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: value,
			Value: value,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetRoleDictionary 获取角色字典
func (*commonService) GetRoleDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	roleMap := foundation.RoleFoundation.WeightByRoleMap
	for k, v := range roleMap {
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: k,
			Value: strconv.Itoa(int(v)),
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetProcessorDictionary 获取处理器字典
func (*commonService) GetProcessorDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constants.Processor)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		processorStructure := val.(constants.ProcessorStructure)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: processorStructure.Name,
			Value: processorStructure.Code,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetAppGenreDictionary 获取应用类型字典
func (*commonService) GetAppGenreDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constants.AppGenre)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		value := val.(string)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: value,
			Value: value,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetLimiterDictionary 获取限流器字典
func (*commonService) GetLimiterDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(limiter.LimiterConfig)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		limiterStructure := val.(entity.Limiter)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label:    limiterStructure.Name,
			Value:    limiterStructure.Code,
			Appendix: limiterStructure.Parameter,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetLimiterModeDictionary 获取限流器模式字典
func (*commonService) GetLimiterModeDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constants.LimiterMode)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		limiterMode := val.(string)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: limiterMode,
			Value: limiterMode,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}

// GetNameListDictionary 获取黑白名单字典
func (*commonService) GetNameListDictionary() (error, []returnee.GetDictionary) {
	var dictionaryList []returnee.GetDictionary
	var valInfo = reflect.ValueOf(constants.NameList)
	for i := 0; i < valInfo.NumField(); i++ {
		val := valInfo.Field(i).Interface()
		limiterMode := val.(string)
		dictionaryList = append(dictionaryList, returnee.GetDictionary{
			Label: limiterMode,
			Value: limiterMode,
		})
	}
	// 如果为空，就返回空数组，而不是nil
	if dictionaryList == nil {
		dictionaryList = []returnee.GetDictionary{}
	}
	return nil, dictionaryList
}
