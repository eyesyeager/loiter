package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/config"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"reflect"
	"strings"
)

/**
 * 处理器业务层
 * @auth eyesYeager
 * @date 2024/1/11 19:07
 */

var ProcessorService = processorService{}

type processorService struct {
}

// SaveAppProcessor 更新应用处理器
func (*processorService) SaveAppProcessor(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.SaveAppProcessor) error {
	// 获取应用当前处理器
	var checkAppProcessorList []entity.AppProcessor
	if err := global.MDB.Where(&entity.AppProcessor{AppId: data.AppId}).Find(&checkAppProcessorList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	appGenreMap := make(map[string]struct{})
	for _, item := range checkAppProcessorList {
		appGenreMap[item.Genre] = struct{}{}
	}

	// 更新应用处理器
	valInfo := reflect.ValueOf(data)
	typeInfo := reflect.TypeOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		fieldType := typeInfo.Field(i)
		if fieldType.Name == "AppId" {
			continue
		}
		processorList := valInfo.FieldByName(fieldType.Name).Interface().([]string)
		_, ok := appGenreMap[strings.ToLower(fieldType.Name)]
		var err error
		if ok && len(processorList) == 0 {
			// 删除操作
			err = global.MDB.Where(&entity.AppProcessor{
				AppId: data.AppId,
				Genre: strings.ToLower(fieldType.Name),
			}).Unscoped().Delete(&entity.AppProcessor{}).Error
		} else if ok && len(processorList) != 0 {
			// 修改操作
			condition := entity.AppProcessor{
				AppId: data.AppId,
				Genre: strings.ToLower(fieldType.Name),
			}
			err = global.MDB.Model(&condition).Where(&condition).Update("codes",
				strings.Join(processorList, config.Program.PluginConfig.ProcessorDelimiter)).Error
		} else if !ok && len(processorList) != 0 {
			// 新增操作
			err = global.MDB.Create(&entity.AppProcessor{
				AppId: data.AppId,
				Genre: strings.ToLower(fieldType.Name),
				Codes: strings.Join(processorList, config.Program.PluginConfig.ProcessorDelimiter),
			}).Error
		}
		if err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}

	// 记录操作日志
	go func() {
		marshal, _ := json.Marshal(data)
		LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.SaveAppProcessor, marshal))
	}()
	return nil
}

// DeleteAppProcessor 删除应用处理器
func (*processorService) DeleteAppProcessor(appId uint) error {
	// 删除处理器配置
	if err := global.MDB.Where(&entity.AppProcessor{AppId: appId}).Unscoped().Delete(&entity.AppProcessor{}).Error; err != nil {
		return err
	}
	// 删除限流器配置
	if err := global.MDB.Where(&entity.AppLimiter{AppId: appId}).Unscoped().Delete(&entity.AppLimiter{}).Error; err != nil {
		return err
	}
	// 删除黑白名单配置
	if err := global.MDB.Where(&entity.AppNameList{AppId: appId}).Unscoped().Delete(&entity.AppNameList{}).Error; err != nil {
		return err
	}
	if err := global.MDB.Where(&entity.NameList{AppId: appId}).Unscoped().Delete(&entity.NameList{}).Error; err != nil {
		return err
	}
	return nil
}

// GetProcessorByPage 分页获取应用处理器
func (*processorService) GetProcessorByPage(data receiver.GetProcessorByPage) (err error, res returnee.GetProcessorByPage) {
	// 获取拼装应用
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	var appIdList []uint
	if err = global.MDB.Raw(`SELECT id FROM app a
          				WHERE 0 = ? OR a.id = ?
						ORDER BY a.created_at DESC
						LIMIT ?, ?`, data.AppId, data.AppId, offset, limit).Scan(&appIdList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	if len(appIdList) == 0 {
		return err, res
	}
	// 拼接应用信息
	var processorPOList []po.GetProcessorByPage
	if err = global.MDB.Raw(`SELECT a.id AppId, a.name, ap.genre, ap.codes 
						FROM app a, app_processor ap 
						WHERE a.id = ap.app_id AND (0 = ? OR a.id = ?) AND a.id IN (?)
						ORDER BY a.created_at DESC`, data.AppId, data.AppId, appIdList).Scan(&processorPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 组装inner
	innerMap := make(map[uint][]po.GetProcessorByPage)
	for _, item := range processorPOList {
		innerMap[item.AppId] = append(innerMap[item.AppId], item)
	}
	var resInnerList []returnee.GetProcessorByPageInner
	for k, v := range innerMap {
		resInner := returnee.GetProcessorByPageInner{AppId: k, AppName: v[0].Name}
		for _, item := range v {
			if item.Genre == constants.Processor.Filter.Code {
				resInner.FilterStr = item.Codes
			} else if item.Genre == constants.Processor.Aid.Code {
				resInner.AidStr = item.Codes
			} else if item.Genre == constants.Processor.Exception.Code {
				resInner.ExceptionStr = item.Codes
			} else if item.Genre == constants.Processor.Final.Code {
				resInner.FinalStr = item.Codes
			}
		}
		resInnerList = append(resInnerList, resInner)
	}
	// 获取处理器信息
	var processorList []entity.Processor
	if err = global.MDB.Find(&processorList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	processorMap := make(map[string]string)
	for _, item := range processorList {
		processorMap[item.Code] = item.Name
	}
	// 翻译处理器名
	for index, item := range resInnerList {
		if item.FilterStr != "" {
			resInnerList[index].Filter, resInnerList[index].FilterCode = translateStrProcessor(processorMap, item.FilterStr)
		}
		if item.AidStr != "" {
			resInnerList[index].Aid, resInnerList[index].AidCode = translateStrProcessor(processorMap, item.AidStr)
		}
		if item.ExceptionStr != "" {
			resInnerList[index].Exception, resInnerList[index].ExceptionCode = translateStrProcessor(processorMap, item.ExceptionStr)
		}
		if item.FinalStr != "" {
			resInnerList[index].Final, resInnerList[index].FinalCode = translateStrProcessor(processorMap, item.FinalStr)
		}
	}
	// 获取总数
	var count int64
	checkApp := entity.App{Model: gorm.Model{ID: data.AppId}}
	if err = global.MDB.Model(&checkApp).Where(checkApp).Count(&count).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 组装返回
	res.PageStruct = data.PageStruct
	res.Total = count
	res.Data = resInnerList
	return err, res
}

// GetProcessorByGenre 根据种类获取处理器
func (*processorService) GetProcessorByGenre(genre string) (error, []returnee.GetDictionary) {
	var processorList []entity.Processor
	if err := global.MDB.Where(&entity.Processor{Genre: genre}).Find(&processorList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil
	}
	var res []returnee.GetDictionary
	for _, item := range processorList {
		res = append(res, returnee.GetDictionary{
			Label: item.Name,
			Value: item.Code,
		})
	}
	return nil, res
}

/***********************************************************************
 *                              help
 ***********************************************************************/

// translateStrProcessor 翻译字符串拼接的处理器
func translateStrProcessor(processorMap map[string]string, codes string) ([]string, []string) {
	var res []string
	var codeRes []string
	split := strings.Split(codes, config.Program.PluginConfig.ProcessorDelimiter)
	for _, item := range split {
		name, ok := processorMap[item]
		if ok {
			res = append(res, name)
		} else {
			res = append(res, item)
		}
		codeRes = append(codeRes, item)
	}
	return res, codeRes
}
