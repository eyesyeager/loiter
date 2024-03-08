package processor

import (
	"errors"
	"fmt"
	"loiter/app/plugin/filter/namelist"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/service"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/model/entity"
	"loiter/model/receiver"
	"loiter/utils"
	"net/http"
	"strings"
)

/**
 * 黑白名单业务层
 * @auth eyesYeager
 * @date 2024/2/29 17:54
 */

type nameListService struct {
}

var NameListService = nameListService{}

// UpdateAppNameList 更新应用黑白名单
func (*nameListService) UpdateAppNameList(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppNameList) error {
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
		if err := global.MDB.Where(&entity.AppNameList{
			AppId: data.AppId,
			Genre: data.Genre,
		}).Unscoped().Delete(&entity.AppNameList{}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, tx.Error.Error()))
		}
	}
	// 打印操作日志
	go service.LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppNameList, data.Genre, constant.GetTurnstileName(data.Turnstile)))
	return nil
}

// AddNameListIp 添加黑白名单ip
func (*nameListService) AddNameListIp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.AddNameListIp) error {
	// 参数合法性校验
	if !namelist.CheckNameListGenre(data.Genre) {
		return errors.New(fmt.Sprintf("非法genre参数：%s", data.Genre))
	}
	// 获取对应应用
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, data.AppId).Error; err != nil {
		return errors.New(fmt.Sprintf("获取id为%d的应用信息失败，请检查应用是否有效或网络是否正常", data.AppId))
	}
	// 构建黑白名单实体
	ipList := strings.Split(data.IpListStr, config.Program.PluginConfig.NameListIpDelimiter)
	var nameListEntityList []entity.NameList
	for _, item := range ipList {
		nameListEntityList = append(nameListEntityList, entity.NameList{
			AppId: data.AppId,
			Genre: data.Genre,
			Ip:    item,
		})
	}
	// 添加ip
	if err := global.MDB.Create(&nameListEntityList).Error; err != nil {
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
	go service.LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.AddNameListIp, checkApp.Name, data.Genre, data.IpListStr))
	return nil
}

// DeleteNameListIp 删除黑白名单ip
func (*nameListService) DeleteNameListIp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.DeleteNameListIp) error {
	// 参数合法性校验
	if !namelist.CheckNameListGenre(data.Genre) {
		return errors.New(fmt.Sprintf("非法genre参数：%s", data.Genre))
	}
	// 获取对应应用
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, data.AppId).Error; err != nil {
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
	go service.LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DeleteNameListIp, checkApp.Name, data.Genre, data.Ip))
	return nil
}
