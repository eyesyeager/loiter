package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/controller/result"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"time"
)

/**
 * 通知业务层
 * @auth eyesYeager
 * @date 2024/2/23 11:21
 */

type noticeService struct {
}

var NoticeService = noticeService{}

// GetNoticeList 分页获取通知列表
func (*noticeService) GetNoticeList(data receiver.GetNoticeList) (err error, res returnee.GetNoticeList) {
	tx := global.MDB.Table("notice").Select("*")
	if data.AppName != "" {
		tx = tx.Where("app_name = ?", data.AppName)
	}
	if data.Title != "" {
		tx = tx.Where("title LIKE ?", "%"+data.Title+"%")
	}
	if data.Genre != "" {
		tx = tx.Where("genre = ?", data.Genre)
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
	var resDataPO []entity.Notice
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&resDataPO).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 时间格式化
	var resData []returnee.GetNoticeListInner
	for _, item := range resDataPO {
		var resItem returnee.GetNoticeListInner
		_ = copier.Copy(&resItem, &item)
		resItem.CreatedAt = item.CreatedAt.Format(time.DateTime)
		// 邮件通知不返回内容
		if resItem.Genre == constants.Notice.Email {
			resItem.Content = ""
		}
		resData = append(resData, resItem)
	}

	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = resData
	return err, res
}

// GetEmailNoticeContent 获取邮件通知内容
func (*noticeService) GetEmailNoticeContent(id uint) (err error, content string) {
	noticeEntity := entity.Notice{
		Model:  gorm.Model{ID: id},
		Genre:  constants.Notice.Email,
		Secret: false,
	}
	if err = global.MDB.Where(&noticeEntity).First(&noticeEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err, "保密通知"
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), content
		}
	}
	return err, noticeEntity.Content
}
