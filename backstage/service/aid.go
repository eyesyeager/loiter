package service

import (
	"errors"
	"fmt"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/backstage/utils"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/receiver"
	"net/http"
	"strings"
)

/**
 * 响应处理器业务层
 * @auth eyesYeager
 * @date 2024/1/26 16:10
 */

type aidService struct {
}

var AidService = aidService{}

// UpdateAppAid 更新应用响应处理器
func (*aidService) UpdateAppAid(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppAid) error {
	// 获取现有配置
	var checkAppAid entity.AppAid
	checkAppAidTX := global.MDB.Where(&entity.AppAid{AppId: data.AppId}).First(&checkAppAid)

	// 如果待更新配置与原配置相通，则取消更新
	newAidName := strings.Join(data.AidNameSlice, config.Program.PluginConfig.AidDelimiter)
	if (checkAppAidTX.RowsAffected == 0 && newAidName == "") || newAidName == checkAppAid.AidName {
		return errors.New("响应处理器更新配置与当前配置相同！")
	}

	// 如果待更新响应处理器为空，则删除原数据
	if len(data.AidNameSlice) == 0 {
		if err := global.MDB.Where(&entity.AppAid{AppId: data.AppId}).Unscoped().Delete(&entity.AppAid{}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
		go LogService.App(r, userClaims.Uid, data.AppId,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppAid, checkAppAid.AidName, newAidName))
		return nil
	}

	// 校验响应处理器
	var aidSlice []entity.Aid
	if err := global.MDB.Where("name IN ?", data.AidNameSlice).Find(&aidSlice).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	if len(aidSlice) != len(data.AidNameSlice) {
		validNameMap := make(map[string]struct{})
		for _, item := range aidSlice {
			validNameMap[item.Name] = struct{}{}
		}
		var invalidNameSlice []string
		for _, item := range data.AidNameSlice {
			if _, ok := validNameMap[item]; !ok {
				invalidNameSlice = append(invalidNameSlice, item)
			}
		}
		return errors.New(fmt.Sprintf("更新响应处理器失败! 存在无效处理器: [%s]", strings.Join(invalidNameSlice, config.Program.PluginConfig.AidDelimiter)))
	}

	// 更新处理器配置
	if checkAppAidTX.RowsAffected == 0 {
		// 如果原配置为空，则插入
		if err := global.MDB.Create(&entity.AppAid{
			AppId:   data.AppId,
			AidName: newAidName,
		}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	} else {
		// 如果原配置不为空，则修改
		if err := global.MDB.Model(&entity.AppAid{}).Where("app_id", data.AppId).Update("aid_name", newAidName).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}

	// 记录操作日志
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppAid, checkAppAid.AidName, newAidName))
	return nil
}
