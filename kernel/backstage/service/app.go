package service

import (
	"zliway/global"
	"zliway/kernel/backstage/model/entity"
	"zliway/kernel/backstage/model/receiver"
)

/**
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

type appService struct {
}

var AppService = appService{}

func (appService *appService) AddApp(data receiver.AppAdd) error {
	app := entity.App{
		Host:    data.Host,
		Pattern: data.Pattern,
		Status:  data.Status,
	}
	err := global.MDB.Create(&app).Error
	if err != nil {
		global.Log.Error("fail to add app！" + err.Error())
		return err
	}
	return err
}
