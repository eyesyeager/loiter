package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/controller/result"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/model/structure"
	"loiter/utils"
	"net/http"
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/9/27 12:36
 */

type logService struct {
}

var LogService = logService{}

// Universal 记录通用日志
func (*logService) Universal(r *http.Request, operatorId uint, logUniversalStruct structure.LogUniversalStruct) {
	// 获取用户信息
	var checkUser = entity.User{Model: gorm.Model{ID: operatorId}}
	if err := global.MDB.First(&checkUser, operatorId).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" operatorId: ", operatorId,
			" error: ", err.Error())
	}
	// 构建日志结构
	logUniversal := entity.LogUniversal{
		Operator: checkUser.Username,
		Ip:       utils.GetIp(r),
		Browser:  utils.GetBrowser(r),
		Title:    logUniversalStruct.Title,
		Content:  logUniversalStruct.Content,
	}

	// 插入数据库
	if err := global.MDB.Create(&logUniversal).Error; err != nil {
		global.AppLogger.Error("general log persistence failed,",
			" operator: ", logUniversal.Operator,
			" title: ", logUniversalStruct.Title,
			" content: ", logUniversalStruct.Content,
			" error: ", err.Error())
	}
}

// App 记录应用操作日志
func (*logService) App(r *http.Request, operatorId uint, appId uint, logUniversalStruct structure.LogUniversalStruct) {
	// 获取用户信息
	var checkUser = entity.User{Model: gorm.Model{ID: operatorId}}
	if err := global.MDB.First(&checkUser, operatorId).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" operatorId: ", operatorId,
			" error: ", err.Error())
	}
	// 获取应用信息
	var checkApp = entity.App{Model: gorm.Model{ID: appId}}
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" appId: ", operatorId,
			" error: ", err.Error())
	}

	// 构建日志结构
	logUniversal := entity.LogUniversal{
		Operator: checkUser.Username,
		Ip:       utils.GetIp(r),
		Browser:  utils.GetBrowser(r),
		Title:    logUniversalStruct.Title,
		Content:  fmt.Sprintf("应用名：%s；%s", checkApp.Name, logUniversalStruct.Content),
	}

	// 插入数据库
	if err := global.MDB.Create(&logUniversal).Error; err != nil {
		global.AppLogger.Error("general log persistence failed,",
			" operator: ", logUniversal.Operator,
			" title: ", logUniversalStruct.Title,
			" content: ", logUniversalStruct.Content,
			" error: ", err.Error())
	}
}

// GetUniversalLog 获取通用日志
func (*logService) GetUniversalLog(data receiver.GetUniversalLog) (err error, res returnee.GetUniversalLog) {
	// 构建请求条件
	tx := global.MDB.Table("log_universal").Select("ip, browser, title, content, created_at, operator")
	if data.OperatorName != "" {
		tx = tx.Where("operator = ?", data.OperatorName)
	}
	if data.Title != "" {
		tx = tx.Where("title LIKE ?", "%"+data.Title+"%")
	}
	if data.Content != "" {
		tx = tx.Where("content LIKE ?", "%"+data.Content+"%")
	}
	if data.TimeBegin != "" {
		data.TimeBegin += " 00:00:00"
		tx = tx.Where("created_at >= ?", data.TimeBegin)
	}
	if data.TimeEnd != "" {
		data.TimeEnd += " 23:59:59"
		tx = tx.Where("created_at <= ?", data.TimeEnd)
	}
	// 查总数
	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 查数据
	var resDataPO []po.GetUniversalLogInner
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&resDataPO).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 时间格式化
	var resData []returnee.GetUniversalLogInner
	for _, item := range resDataPO {
		var resItem returnee.GetUniversalLogInner
		_ = copier.Copy(&resItem, &item)
		resItem.CreatedAt = item.CreatedAt.Format(time.DateTime)
		resData = append(resData, resItem)
	}

	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = resData
	return err, res
}
