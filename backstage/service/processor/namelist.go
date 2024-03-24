package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"loiter/app/plugin/filter/namelist"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/service"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/model/entity"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
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
func (n *nameListService) UpdateAppNameList(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppNameList) error {
	// 查看当前配置是否存在
	var checkAppNameList []entity.AppNameList
	if err := global.MDB.Where(&entity.AppNameList{AppId: data.AppId}).Find(&checkAppNameList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 可行性校验
	if err := n.updateGenreNameList(data.AppId, constants.NameList.Black, data.Black, checkAppNameList); err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if err := n.updateGenreNameList(data.AppId, constants.NameList.White, data.White, checkAppNameList); err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 打印操作日志
	go func() {
		marshal, _ := json.Marshal(data)
		service.LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppNameList, marshal))
	}()
	return nil
}

// GetAppNameList 获取应用黑白名单状态
func (*nameListService) GetAppNameList(appId uint) (err error, res returnee.GetAppNameList) {
	var appNameList []entity.AppNameList
	if err = global.MDB.Where(&entity.AppNameList{AppId: appId}).Find(&appNameList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for _, item := range appNameList {
		if item.Genre == constants.NameList.Black {
			res.Black = true
		} else {
			res.White = true
		}
	}
	return nil, res
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
	var nameListEntityList []entity.NameList
	for _, item := range data.IpList {
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
	if data.Genre == constants.NameList.Black {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == constants.NameList.White {
		iNameList, ok = container.WhiteNameListByAppMap[checkApp.Host]
	}
	if ok {
		if err := iNameList.Refresh(); err != nil {
			global.AppLogger.Error(fmt.Sprintf("failed to refresh bloom filter! genre: %s, host: %s, error: %s", data.Genre, checkApp.Host, err.Error()))
		}
	}
	// 打印操作日志
	go func() {
		marshal, _ := json.Marshal(data.IpList)
		service.LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.AddNameListIp, checkApp.Name, data.Genre, marshal))
	}()
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
	if data.Genre == constants.NameList.Black {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == constants.NameList.White {
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

// updateGenreNameList 更新应用单个类型的名单状态
func (*nameListService) updateGenreNameList(appId uint, genre string, status bool, checkAppNameList []entity.AppNameList) error {
	// 判断该类型是否已经存在
	alreadyExist := false
	for _, item := range checkAppNameList {
		if item.Genre == genre {
			alreadyExist = true
			continue
		}
	}
	// 如果两者状态相同，则无需操作
	if status == alreadyExist {
		return nil
	}
	if status {
		// 插入操作
		if err := global.MDB.Create(&entity.AppNameList{
			AppId: appId,
			Genre: genre,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	} else {
		// 删除操作
		if err := global.MDB.Where(&entity.AppNameList{
			AppId: appId,
			Genre: genre,
		}).Unscoped().Delete(&entity.AppNameList{}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	return nil
}
