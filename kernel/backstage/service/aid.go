package service

import (
	"errors"
	"fmt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/utils"
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
	newAidName := strings.Join(data.AidNameSlice, config.Program.AidDelimiter)
	if (checkAppAidTX.RowsAffected == 0 && newAidName == "") || newAidName == checkAppAid.AidName {
		return errors.New("响应处理器更新配置与当前配置相同！")
	}

	// 如果待更新响应处理器为空，则删除原数据
	if len(data.AidNameSlice) == 0 {
		if err := global.MDB.Where(&entity.AppAid{AppId: data.AppId}).Unscoped().Delete(&entity.AppAid{}).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppAid()-Delete", err.Error()))
		}
		go func() {
			//var app entity.App
			//if err := global.MDB.First(&app, data.AppId).Error; err != nil {
			//	global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-doLog-insert", err.Error()))
			//}
			//LogService.Universal(r, userClaims.Uid,
			//	constant.BuildUniversalLog(constant.LogUniversal.UpdateAppPassageway, app.Name, "", newPassagewayName))
		}()
	}

	// 校验响应处理器

	// 如果原配置为空，则插入

	// 如果原配置不为空，则更新

	// 记录操作日志
	return nil
}
