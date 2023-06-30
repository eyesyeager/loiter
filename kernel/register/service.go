package register

import (
	"strconv"
	"zliway/global"
	"zliway/kernel/backstage/model/entity"
	"zliway/kernel/register/constant"
)

/**
 * 业务层
 * @author eyesYeager
 * @date 2023/4/27 10:52
 */

func BuildAppSingleton() error {
	// 获取app切片
	var appSlice []entity.App
	if err := global.MDB.Where(entity.App{
		Pattern: constant.AppPattern["singleton"],
		Status:  constant.AppStatus["normal"],
	}).Find(&appSlice).Error; err != nil {
		global.Log.Error("failed to obtain app slices in singleton mode with normal status: " + err.Error())
		return err
	}
	if len(appSlice) == 0 {
		// 清空map
		AppSingletonHolder = make(map[string][]ServerModel)
		return nil
	}

	// 根据app切片获取对应server
	appIds := ""
	for _, item := range appSlice {
		appIds += "," + strconv.Itoa(int(item.Id))
	}
	var serverSlice []serverModelWithAppId
	if err := global.MDB.Raw("select app.app, server.server, server.weight, server.group from app, basket, server where app.id=basket.app_id and basket.app_id in (?) and basket.id=server.basket_id and server.status=?", appIds[1:], constant.ServerStatus["normal"]).Scan(&serverSlice).Error; err != nil {
		global.Log.Error("failed to obtain server slices: " + err.Error())
		return err
	}
	if len(serverSlice) == 0 {
		// 清空map
		AppSingletonHolder = make(map[string][]ServerModel)
		return nil
	}

	// 组装结果
	appSingletonHolderTemp := map[string][]ServerModel{}
	for _, appItem := range appSlice {
		for _, serverItem := range serverSlice {
			if _, ok := appSingletonHolderTemp[serverItem.App]; ok {

			} else {

			}
		}
	}
	AppSingletonHolder = appSingletonHolderTemp
	return nil
}

// -------------------------------------- 局部结构体 ------------------------------------------

type serverModelWithAppId struct {
	App    string
	Server string
	Weight uint
	Group  string
}
