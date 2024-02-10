package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/controller/result"
	"loiter/backstage/utils"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/po"
	"loiter/kernel/model/receiver"
	"loiter/kernel/model/returnee"
	"loiter/kernel/model/structure"
	"net/http"
	"strconv"
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/9/27 12:36
 */

type logService struct {
}

var LogService = logService{}

// Login 记录登录日志
func (*logService) Login(r *http.Request, uid uint) {
	logLogin := entity.LogLogin{
		Uid:     uid,
		Ip:      helper.GetIp(r),
		Browser: helper.GetBrowser(r),
	}

	// 插入数据库
	if err := global.MDB.Create(&logLogin).Error; err != nil {
		global.AppLogger.Error("User login log insertion failed for userId ", strconv.Itoa(int(uid)), ", error:", err.Error())
	}
}

// Universal 记录通用日志
func (*logService) Universal(r *http.Request, operatorId uint, logUniversalStruct structure.LogUniversalStruct) {
	// 获取用户信息
	var checkUser = entity.User{Model: gorm.Model{ID: operatorId}}
	if err := global.MDB.First(&checkUser).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" operatorId: ", operatorId,
			" error: ", err.Error())
	}
	// 构建日志结构
	logUniversal := entity.LogUniversal{
		Operator: checkUser.Username,
		Ip:       helper.GetIp(r),
		Browser:  helper.GetBrowser(r),
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
	if err := global.MDB.First(&checkUser).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" operatorId: ", operatorId,
			" error: ", err.Error())
	}
	// 获取应用信息
	var checkApp = entity.App{Model: gorm.Model{ID: appId}}
	if err := global.MDB.First(&checkApp).Error; err != nil {
		global.AppLogger.Error("failed to obtain user information,",
			" appId: ", operatorId,
			" error: ", err.Error())
	}

	// 构建日志结构
	logUniversal := entity.LogUniversal{
		Operator: checkUser.Username,
		Ip:       helper.GetIp(r),
		Browser:  helper.GetBrowser(r),
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

// GetLoginLog 获取登录日志
func (*logService) GetLoginLog(data receiver.GetLoginLog) (err error, res returnee.GetLoginLog) {
	// 构建请求条件
	tx := global.MDB.Table("log_login ll").Select("ll.ip, ll.browser, ll.created_at, u.username").Joins("INNER JOIN user u on u.id = ll.uid")
	if data.Username != "" {
		tx = tx.Where("u.username = ?", data.Username)
	}
	if data.LoginTimeBegin != "" {
		tx = tx.Where("ll.created_at >= ?", data.LoginTimeBegin)
	}
	if data.LoginTimeEnd != "" {
		tx = tx.Where("ll.created_at <= ?", data.LoginTimeEnd)
	}
	// 查总数
	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 查数据
	var resDataPO []po.GetLoginLogInner
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("ll.created_at DESC").Limit(limit).Offset(offset).Find(&resDataPO).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 时间格式化
	var resData []returnee.GetLoginLogInner
	for _, item := range resDataPO {
		var resItem returnee.GetLoginLogInner
		_ = copier.Copy(&resItem, &item)
		resItem.CreatedAt = item.CreatedAt.Format(time.DateTime)
		resData = append(resData, resItem)
	}

	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = resData
	return err, res
}

// GetUniversalLog 获取通用日志
func (*logService) GetUniversalLog(data receiver.GetUniversalLog) (err error, res returnee.GetUniversalLog) {
	// 构建请求条件
	tx := global.MDB.Table("log_universal lu").Select("lu.ip, lu.browser, lu.title, lu.content, lu.created_at, u.username operator").Joins("INNER JOIN user u on u.id = lu.operator_id")
	if data.OperatorName != "" {
		tx = tx.Where("u.username = ?", data.OperatorName)
	}
	if data.Title != "" {
		tx = tx.Where("lu.title LIKE ?", "%"+data.Title+"%")
	}
	if data.Content != "" {
		tx = tx.Where("lu.content LIKE ?", "%"+data.Content+"%")
	}
	if data.LoginTimeBegin != "" {
		tx = tx.Where("lu.created_at >= ?", data.LoginTimeBegin)
	}
	if data.LoginTimeEnd != "" {
		tx = tx.Where("lu.created_at <= ?", data.LoginTimeEnd)
	}
	// 查总数
	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 查数据
	var resDataPO []po.GetUniversalLogInner
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("lu.created_at DESC").Limit(limit).Offset(offset).Find(&resDataPO).Error; err != nil {
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