package service

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/utils"
	"loiter/kernel/model/entity"
	"loiter/kernel/model/receiver"
	"net/http"
	"strings"
)

/**
 * 通道业务层
 * @auth eyesYeager
 * @date 2024/1/11 19:07
 */

type passagewayService struct {
}

var PassagewayService = passagewayService{}

// UpdateAppPassageway 更新应用通道
func (*passagewayService) UpdateAppPassageway(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppPassageway) error {
	// 获取应用当前通道
	var checkAppPassageway entity.AppPassageway
	checkAppPassagewayTX := global.MDB.Where(&entity.AppPassageway{AppId: data.AppID}).First(&checkAppPassageway)

	// 如果新通道与当前通道相同，则取消更新
	newPassagewayName := strings.Join(data.PassagewayNameSlice, ",")
	if checkAppPassageway.PassagewayName == newPassagewayName {
		return errors.New("通道更新配置与当前配置相同！")
	}

	// 校验是否存在对应通道(若更新通道为空，则跳过校验)
	var passagewaySlice []entity.Passageway
	if len(data.PassagewayNameSlice) != 0 {
		if err := global.MDB.Where("name IN ?", data.PassagewayNameSlice).Find(&passagewaySlice).Error; err != nil {
			errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-passagewaySlice", err.Error())
			global.BackstageLogger.Error(errMsg)
			return errors.New(errMsg)
		}
	}
	// 存在非法通道，返回异常
	if len(passagewaySlice) != len(data.PassagewayNameSlice) {
		validNameMap := make(map[string]struct{})
		for _, item := range passagewaySlice {
			validNameMap[item.Name] = struct{}{}
		}
		var invalidNameSlice []string
		for _, item := range data.PassagewayNameSlice {
			if _, ok := validNameMap[item]; !ok {
				invalidNameSlice = append(invalidNameSlice, item)
			}
		}
		return errors.New(fmt.Sprintf("Update passageway failed! There is an illegal passageway name: [%s]", strings.Join(invalidNameSlice, ",")))
	}

	// 如果原应用未配置策略，则直接插入
	if checkAppPassagewayTX.RowsAffected == 0 {
		// 记录操作日志
		passagewayEntity := entity.AppPassageway{
			AppId:          data.AppID,
			PassagewayName: newPassagewayName,
		}
		if err := global.MDB.Create(&passagewayEntity).Error; err != nil {
			errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-Create", err.Error())
			global.BackstageLogger.Error(errMsg)
			return errors.New(errMsg)
		}
		// 记录操作日志
		go func() {
			var app entity.App
			if err := global.MDB.First(&app, data.AppID).Error; err != nil {
				global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-doLog-insert", err.Error()))
			}
			LogService.Universal(r, userClaims.Uid,
				constant.BuildUniversalLog(constant.LogUniversal.UpdateAppPassageway, app.Name, "", newPassagewayName))
		}()
		return nil
	}

	// 更新应用策略
	if err := global.MDB.Model(&entity.AppPassageway{}).Where("app_id", data.AppID).Update("passageway_name", data.PassagewayNameSlice).Error; err != nil {
		errMsg := fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-Update", err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}

	// 记录操作日志
	go func() {
		var app entity.App
		if err := global.MDB.First(&app, data.AppID).Error; err != nil {
			global.BackstageLogger.Error(fmt.Sprintf(result.CommonInfo.DbOperateError, "UpdateAppPassageway()-doLog-update", err.Error()))
		}
		oldPassagewayName := strings.Join(data.PassagewayNameSlice, ",")
		LogService.Universal(r, userClaims.Uid,
			constant.BuildUniversalLog(constant.LogUniversal.UpdateAppPassageway, app.Name, oldPassagewayName, newPassagewayName))
	}()
	return nil
}
