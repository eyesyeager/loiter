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
	"loiter/model/po"
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
	// 构建请求条件
	tx := global.MDB.Table("notice n").Select("n.id, a.name AppName, n.host, n.title, n.genre, n.content, n.remarks, n.created_at").Joins(
		"LEFT JOIN app a on n.host = a.host")
	if data.AppId != 0 {
		tx = tx.Where("a.id = ?", data.AppId)
	}
	if data.Title != "" {
		tx = tx.Where("n.title LIKE ?", "%"+data.Title+"%")
	}
	if data.Genre != "" {
		tx = tx.Where("n.genre = ?", data.Genre)
	}
	if data.TimeBegin != "" {
		data.TimeBegin += " 00:00:00"
		tx = tx.Where("n.created_at >= ?", data.TimeBegin)
	}
	if data.TimeEnd != "" {
		data.TimeEnd += " 23:59:59"
		tx = tx.Where("n.created_at <= ?", data.TimeEnd)
	}

	// 查总数
	var total int64
	if err = tx.Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 查数据
	var resPOList []po.GetNoticeList
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = tx.Order("n.created_at DESC").Limit(limit).Offset(offset).Find(&resPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 数据组装
	var resData []returnee.GetNoticeListInner
	for _, item := range resPOList {
		var resItem returnee.GetNoticeListInner
		_ = copier.Copy(&resItem, &item)
		// 时间格式化
		resItem.CreatedAt = item.CreatedAt.Format(time.DateTime)
		// 邮件通知不返回内容
		if resItem.Genre == constants.Notice.Email {
			resItem.Content = ""
		}
		// appName处理
		if resItem.AppName == "" {
			resItem.AppName = item.Host
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
