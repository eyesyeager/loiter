package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/app/plugin/filter/namelist"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/service"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"time"
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
	if err := n.updateGenreNameList(data.AppId, constants.NameList.Black.Value, data.Black, checkAppNameList); err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if err := n.updateGenreNameList(data.AppId, constants.NameList.White.Value, data.White, checkAppNameList); err != nil {
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

// GetNameList 获取应用黑白名单
func (*nameListService) GetNameList(data receiver.GetNameList) (err error, res returnee.GetNameList) {
	// 构建请求条件
	tx := global.MDB.Table("name_list nl").Select(" nl.id, a.name AppName, a.id AppId, nl.ip, nl.remarks, nl.created_at").Joins(
		"LEFT JOIN app a on nl.app_id = a.id")
	if data.AppId != 0 {
		tx = tx.Where("a.id = ?", data.AppId)
	}
	if data.Genre != "" {
		tx = tx.Where("nl.genre = ?", data.Genre)
	}
	if data.Ip != "" {
		tx = tx.Where("nl.ip = ?", data.Ip)
	}
	if data.Remarks != "" {
		tx = tx.Where("nl.remarks LIKE ?", "%"+data.Remarks+"%")
	}
	if data.TimeBegin != "" {
		data.TimeBegin += " 00:00:00"
		tx = tx.Where("nl.created_at >= ?", data.TimeBegin)
	}
	if data.TimeEnd != "" {
		data.TimeEnd += " 23:59:59"
		tx = tx.Where("nl.created_at <= ?", data.TimeEnd)
	}

	// 查总数
	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 查数据
	var resInnerPOList []po.GetNameList
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("nl.created_at DESC").Limit(limit).Offset(offset).Find(&resInnerPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 数据组装
	var resData []returnee.GetNameListInner
	for _, item := range resInnerPOList {
		var resItem returnee.GetNameListInner
		_ = copier.Copy(&resItem, &item)
		// 时间格式化
		resItem.CreatedAt = item.CreatedAt.Format(time.DateTime)
		resData = append(resData, resItem)
	}

	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = resData
	return err, res
}

// GetAppNameListStatus 获取应用黑白名单状态
func (*nameListService) GetAppNameListStatus(appId uint) (err error, res returnee.GetAppNameList) {
	var appNameList []entity.AppNameList
	if err = global.MDB.Where(&entity.AppNameList{AppId: appId}).Find(&appNameList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for _, item := range appNameList {
		if item.Genre == constants.NameList.Black.Value {
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
			AppId:   data.AppId,
			Genre:   data.Genre,
			Ip:      item,
			Remarks: data.Remarks,
		})
	}
	// 添加ip
	if err := global.MDB.Create(&nameListEntityList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 刷新布隆过滤器
	var iNameList namelist.INameList
	var ok bool
	if data.Genre == constants.NameList.Black.Value {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == constants.NameList.White.Value {
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
		Model: gorm.Model{ID: data.Id},
		AppId: data.AppId,
		Genre: data.Genre,
		Ip:    data.Ip,
	}).Unscoped().Delete(&entity.NameList{}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 刷新布隆过滤器
	var iNameList namelist.INameList
	var ok bool
	if data.Genre == constants.NameList.Black.Value {
		iNameList, ok = container.BlackNameListByAppMap[checkApp.Host]
	}
	if data.Genre == constants.NameList.White.Value {
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
	// 刷新容器
	if err := container.RefreshNameList(appId); err != nil {
		global.AppLogger.Error(err.Error())
	}
	return nil
}
