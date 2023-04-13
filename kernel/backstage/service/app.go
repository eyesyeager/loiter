package service

import (
	"github.com/jinzhu/copier"
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
	// 构建实体
	var app entity.App
	err := copier.Copy(&app, &data)
	if err != nil {
		return err
	}

	// 执行插入
	err = global.MDB.Create(&app).Error
	if err != nil {
		global.Log.Error("fail to add app！" + err.Error())
		return err
	}
	return err
}
