package service

import (
	"errors"
	"fmt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/utils"
	"net/http"
	"strconv"
)

/**
 * 应用实例业务层
 * @auth eyesYeager
 * @date 2024/1/5 14:08
 */

type serverService struct {
}

var ServerService = serverService{}

// AddServer 注册应用实例
func (*serverService) AddServer(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.AddServer) error {
	// 检查应用信息
	checkApp := entity.App{}
	if err := global.MDB.First(&checkApp, data.AppId).Error; err != nil {
		return errors.New(fmt.Sprintf("ID为%s的应用不存在！", strconv.Itoa(int(data.AppId))))
	}
	// 检查应用实例的唯一性
	checkServer := entity.Server{}
	if tx := global.MDB.Where(&entity.Server{
		AppId:   data.AppId,
		Address: data.Address,
	}).First(&checkServer); tx.RowsAffected != 0 {
		return errors.New(fmt.Sprintf("地址为'%s'的实例已存在，实例名为'%s'", data.Address, checkServer.Name))
	}

	// 参数处理
	if data.Name == "" {
		data.Name = utils.GenerateRandString(config.Program.ServerDefaultNameLen)
	}

	// 插入应用实例
	if err := global.MDB.Create(&entity.Server{
		AppId:   data.AppId,
		Name:    data.Name,
		Address: data.Address,
		Remarks: data.Remarks,
	}).Error; err != nil {
		errMsg := fmt.Sprintf(result.ResultInfo.DbOperateError, err.Error())
		global.BackstageLogger.Error(errMsg)
		return errors.New(errMsg)
	}

	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.AddServer, checkApp.Name, data.Name, data.Address, data.Remarks))
	return nil
}
