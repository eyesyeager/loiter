package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/utils"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/receiver"
	"loiter/plugin/filter/namelist"
	"net/http"
	"strings"
)

/**
 * 过滤器业务层
 * @auth eyesYeager
 * @date 2024/1/11 19:07
 */

type filterService struct {
}

var FilterService = filterService{}

// UpdateAppFilter 更新应用过滤器
func (*filterService) UpdateAppFilter(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppFilter) error {
	// 获取应用当前过滤器
	var checkAppFilter entity.AppFilter
	checkAppFilterTX := global.MDB.Where(&entity.AppFilter{AppId: data.AppId}).First(&checkAppFilter)

	// 如果新过滤器与当前过滤器相同，则取消更新
	newFilterName := strings.Join(data.FilterNameSlice, config.Program.PluginConfig.FilterDelimiter)
	if checkAppFilter.FilterName == newFilterName {
		return errors.New("过滤器更新配置与当前配置相同！")
	}

	// 校验是否存在对应过滤器(若更新过滤器为空，则跳过校验)
	var filterSlice []entity.Filter
	if len(data.FilterNameSlice) != 0 {
		if err := global.MDB.Where("name IN ?", data.FilterNameSlice).Find(&filterSlice).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 存在非法过滤器，返回异常
	if len(filterSlice) != len(data.FilterNameSlice) {
		validNameMap := make(map[string]struct{})
		for _, item := range filterSlice {
			validNameMap[item.Name] = struct{}{}
		}
		var invalidNameSlice []string
		for _, item := range data.FilterNameSlice {
			if _, ok := validNameMap[item]; !ok {
				invalidNameSlice = append(invalidNameSlice, item)
			}
		}
		return errors.New(fmt.Sprintf("更新过滤器失败！存在无效过滤器: [%s]", strings.Join(invalidNameSlice, config.Program.PluginConfig.FilterDelimiter)))
	}

	// 如果原应用未配置策略，则直接插入
	if checkAppFilterTX.RowsAffected == 0 {
		// 记录操作日志
		filterEntity := entity.AppFilter{
			AppId:      data.AppId,
			FilterName: newFilterName,
		}
		if err := global.MDB.Create(&filterEntity).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
		// 记录操作日志
		go LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppFilter, "", newFilterName))
		return nil
	}

	// 更新应用策略
	if err := global.MDB.Model(&entity.AppFilter{}).Where("app_id", data.AppId).Update("filter_name", data.FilterNameSlice).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}

	// 记录操作日志
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppFilter, checkAppFilter, newFilterName))
	return nil
}

// UpdateAppLimiter 更新应用限流器
func (*filterService) UpdateAppLimiter(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppLimiter) error {
	// 校验传入限流器的合法性
	var count int64
	var checkLimiter = entity.Limiter{Name: data.LimiterName}
	if err := global.MDB.Model(&checkLimiter).Where(&checkLimiter).Count(&count).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if count == 0 {
		return errors.New(fmt.Sprintf("不存在名字为%s的限流器", data.LimiterName))
	}
	// 获取应用当前所用限流器
	var checkAppLimiter entity.AppLimiter
	checkAppLimiterTX := global.MDB.Where(&entity.AppLimiter{AppId: data.AppId}).First(&checkAppLimiter)
	// 如果原先没有配置限流器，则插入
	if checkAppLimiterTX.RowsAffected == 0 {
		if err := global.MDB.Create(&entity.AppLimiter{
			AppId:       data.AppId,
			LimiterName: data.LimiterName,
			Parameter:   data.Parameter,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
		// 记录操作日志
		go LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppLimiter, "", "", data.LimiterName, data.Parameter))
		return nil
	}

	// 如果原先已经配置了限流器，则修改
	if err := global.MDB.Model(&entity.AppLimiter{}).Where("app_id", data.AppId).Updates(entity.AppLimiter{
		LimiterName: data.LimiterName,
		Parameter:   data.Parameter,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}

	// 记录操作日志
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppLimiter, checkAppLimiter.LimiterName, checkAppLimiter.Parameter, data.LimiterName, data.Parameter))
	return nil
}

// UpdateAppNameList 更新应用黑白名单
func (*filterService) UpdateAppNameList(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppNameList) error {
	// 参数合法性校验
	if !namelist.CheckNameListGenre(data.Genre) {
		return errors.New(fmt.Sprintf("非法genre参数：%s", data.Genre))
	}
	if !constant.CheckTurnstile(data.Turnstile) {
		return errors.New(fmt.Sprintf("非法turnstile参数：%s", data.Genre))
	}
	// 查看当前配置是否存在
	var checkAppNameList = entity.AppNameList{
		AppId: data.AppId,
		Genre: data.Genre,
	}
	tx := global.MDB.Limit(1).Find(&checkAppNameList)
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
	}
	// 可行性校验
	if data.Turnstile == constant.Turnstile.Open {
		if tx.RowsAffected != 0 {
			return errors.New("操作失败，该名单已处于开启状态")
		}
		// 插入配置
		if err := global.MDB.Create(&entity.AppNameList{
			AppId: data.AppId,
			Genre: data.Genre,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
		}
	} else if data.Turnstile == constant.Turnstile.Close {
		if tx.RowsAffected == 0 {
			return errors.New("操作失败，该名单已处于关闭状态")
		}
		// 删除配置
		if err := global.MDB.Where(entity.AppNameList{
			AppId: data.AppId,
			Genre: data.Genre,
		}).Unscoped().Delete(&entity.AppNameList{}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
		}
	}
	// 打印操作日志
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppNameList, data.Genre, constant.GetTurnstileName(data.Turnstile)))
	return nil
}

// AddNameListIp 添加黑白名单ip
func (*filterService) AddNameListIp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.AddNameListIp) error {
	// 参数合法性校验
	if !namelist.CheckNameListGenre(data.Genre) {
		return errors.New(fmt.Sprintf("非法genre参数：%s", data.Genre))
	}
	// 获取对应应用
	var checkApp = entity.App{Model: gorm.Model{ID: data.AppId}}
	if err := global.MDB.First(&checkApp).Error; err != nil {
		return errors.New(fmt.Sprintf("获取id为%d的应用信息失败，请检查应用是否有效或网络是否正常", data.AppId))
	}
	// 构建黑白名单实体
	ipSlice := strings.Split(data.IpSliceStr, config.Program.PluginConfig.NameListIpDelimiter)
	var nameListEntitySlice []entity.NameList
	for _, item := range ipSlice {
		nameListEntitySlice = append(nameListEntitySlice, entity.NameList{
			AppId: data.AppId,
			Genre: data.Genre,
			Ip:    item,
		})
	}
	// 添加ip
	if err := global.MDB.Create(&nameListEntitySlice).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 刷新布隆过滤器
	var iNameList namelist.INameList
	var ok bool
	if data.Genre == namelist.BlackList {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == namelist.WhiteList {
		iNameList, ok = container.WhiteNameListByAppMap[checkApp.Host]
	}
	if ok {
		if err := iNameList.Refresh(); err != nil {
			global.AppLogger.Error(fmt.Sprintf("failed to refresh bloom filter! genre: %s, host: %s, error: %s", data.Genre, checkApp.Host, err.Error()))
		}
	}
	// 打印操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.AddNameListIp, checkApp.Name, data.Genre, data.IpSliceStr))
	return nil
}

// DeleteNameListIp 删除黑白名单ip
func (*filterService) DeleteNameListIp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.DeleteNameListIp) error {
	// 参数合法性校验
	if !namelist.CheckNameListGenre(data.Genre) {
		return errors.New(fmt.Sprintf("非法genre参数：%s", data.Genre))
	}
	// 获取对应应用
	var checkApp = entity.App{Model: gorm.Model{ID: data.AppId}}
	if err := global.MDB.First(&checkApp).Error; err != nil {
		return errors.New(fmt.Sprintf("获取id为%d的应用信息失败，请检查应用是否有效或网络是否正常", data.AppId))
	}
	// 删除ip
	if err := global.MDB.Where(&entity.NameList{
		AppId: data.AppId,
		Genre: data.Genre,
		Ip:    data.Ip,
	}).Unscoped().Delete(&entity.NameList{}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 刷新布隆过滤器
	var iNameList namelist.INameList
	var ok bool
	if data.Genre == namelist.BlackList {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == namelist.WhiteList {
		iNameList, ok = container.WhiteNameListByAppMap[checkApp.Host]
	}
	if ok {
		if err := iNameList.Refresh(); err != nil {
			global.AppLogger.Error(fmt.Sprintf("failed to refresh bloom filter! genre: %s, host: %s, error: %s", data.Genre, checkApp.Host, err.Error()))
		}
	}
	// 打印操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DeleteNameListIp, checkApp.Name, data.Genre, data.Ip))
	return nil
}
