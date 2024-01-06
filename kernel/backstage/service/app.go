package service

import (
	"errors"
	"fmt"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/utils"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/9/26 15:33
 */
type appService struct {
}

var AppService = appService{}

// AddApp 注册应用
func (*appService) AddApp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.AddApp) error {
	// 检查Host的唯一性
	checkApp := entity.App{}
	if tx := global.MDB.Where(&entity.App{Host: data.Host}).First(&checkApp); tx.RowsAffected != 0 {
		return errors.New(fmt.Sprintf("host为'%s'的应用已存在，应用名为'%s'", data.Host, checkApp.Name))
	}
	// 插入应用
	if err := global.MDB.Create(&entity.App{
		Name:    data.Name,
		Host:    data.Host,
		Remarks: data.Remarks,
	}).Error; err != nil {
		errMsg := fmt.Sprintf(result.ResultInfo.DbOperateError, err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}
	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.AddApp, data.Name, data.Host, data.Remarks))
	return nil
}
