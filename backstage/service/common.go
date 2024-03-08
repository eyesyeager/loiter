package service

import (
	"loiter/backstage/constant"
	"loiter/backstage/foundation"
	"loiter/constants"
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
	return nil, dictionaryList
}
