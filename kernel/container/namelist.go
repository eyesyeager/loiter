package container

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/po"
	"loiter/plugin/passageway/filter/namelist"
)

/**
 * 黑白名单容器
 * @auth eyesYeager
 * @date 2024/1/24 11:14
 */

// NameListByAppMap 黑白名单开启类型 by AppHost
var NameListByAppMap map[string][]string

// BlackNameListByAppMap 黑名单结构体 by AppHost
var BlackNameListByAppMap map[string]namelist.INameList

// WhiteNameListByAppMap 白名单结构体 by AppHost
var WhiteNameListByAppMap map[string]namelist.INameList

// InitNameList 初始化黑白名单容器
func InitNameList() {
	global.AppLogger.Info("start initializing the NameList container")
	// 获取有效应用黑白名单配置
	var appNameListSlice []po.GetAppNameList
	if affected := global.MDB.Raw(`SELECT a.host, anl.genre
					FROM app a, app_name_list anl
					WHERE a.status = ? AND a.id = anl.app_id`, constant.Status.Normal).Scan(&appNameListSlice).RowsAffected; affected == 0 {
		global.AppLogger.Info("there is currently no valid NameList configuration")
		return
	}

	// 构建容器
	var existMap = make(map[string]struct{})
	var genreContainerMap = make(map[string][]string)           // 开启类型临时容器
	var blackContainerMap = make(map[string]namelist.INameList) // 黑名单临时容器
	var whiteContainerMap = make(map[string]namelist.INameList) // 白名单临时容器
	for _, item := range appNameListSlice {
		// 开启类型临时容器
		if _, ok := existMap[item.Host]; ok {
			genreContainerMap[item.Host] = append(genreContainerMap[item.Host], item.Genre)
		} else {
			existMap[item.Host] = struct{}{}
			genreContainerMap[item.Host] = []string{item.Genre}
		}
		// 黑白名单临时容器
		err, nameList := namelist.NewNameList(item.Host, item.Genre)
		if err != nil {
			global.AppLogger.Error(fmt.Sprintf("nameList container creation failed! host: %s, genre: %s, error: %s", item.Host, item.Genre, err.Error()))
			continue
		}
		if item.Genre == namelist.BlackList {
			blackContainerMap[item.Host] = nameList
		}
		if item.Genre == namelist.WhiteList {
			whiteContainerMap[item.Host] = nameList
		}
	}
	NameListByAppMap = genreContainerMap
	BlackNameListByAppMap = blackContainerMap
	WhiteNameListByAppMap = whiteContainerMap
	global.AppLogger.Info("complete the initialization of NameList container")
}

// RefreshNameList 刷新黑白名单容器
func RefreshNameList(appId uint) error {
	global.AppLogger.Info(fmt.Sprintf("start refreshing the NameList container under the application with appId %d", appId))
	// 获取有效黑白名单容器
	var appNameListSlice []po.GetAppNameList
	tx := global.MDB.Raw(`SELECT a.host, anl.genre
						FROM app a, app_name_list anl
						WHERE a.id = ? AND a.status = ? AND a.id = anl.app_id`, appId, constant.Status.Normal).Scan(&appNameListSlice)
	// 查询错误则返回错误信息
	if tx.Error != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "RefreshNameList()-Raw", tx.Error.Error()))
	}
	// 查询为空则删除容器元素
	if tx.RowsAffected == 0 {
		var checkApp = entity.App{Model: gorm.Model{ID: appId}}
		if err := global.MDB.First(&checkApp).Error; err != nil {
			return errors.New(fmt.Sprintf("appId为%d的应用不存在或者无效！", appId))
		}
		delete(NameListByAppMap, checkApp.Host)
		return nil
	}
	// 刷新容器
	var newNameList []string
	for _, item := range appNameListSlice {
		newNameList = append(newNameList, item.Genre)
		err, nameList := namelist.NewNameList(item.Host, item.Genre)
		if err != nil {
			return errors.New(fmt.Sprintf("黑白名单实例创建失败! host: %s，名单类型：%s，错误信息：%s", item.Host, item.Genre, err.Error()))
		}
		if item.Genre == namelist.BlackList {
			BlackNameListByAppMap[item.Host] = nameList
		}
		if item.Genre == namelist.WhiteList {
			WhiteNameListByAppMap[item.Host] = nameList
		}
	}
	NameListByAppMap[appNameListSlice[0].Host] = newNameList
	global.AppLogger.Info(fmt.Sprintf("complete the refresh of NameList container under the application with appId %d", appId))
	return nil
}
