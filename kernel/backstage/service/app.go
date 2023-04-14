package service

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"zliway/global"
	"zliway/kernel/backstage/model/entity"
	"zliway/kernel/backstage/model/receiver"
	"zliway/kernel/backstage/model/returnee"
)

/**
 * 应用相关业务
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

type appService struct {
}

var AppService = appService{}

// AddApp 添加应用
func (appService *appService) AddApp(r *http.Request, data receiver.AppAdd) error {
	// 检查应用是否已经注册
	var currentApp entity.App
	if tx := global.MDB.Where(entity.App{App: data.App}).Find(&currentApp); tx.RowsAffected != 0 {
		return errors.New("the application '" + data.App + "' has already been registered with ID " + strconv.Itoa(int(currentApp.Id)))
	}

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

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addApp"], string(remarks))
	return err
}

// AddServerApp 添加Server
func (appService *appService) AddServerApp(r *http.Request, data receiver.AppServerAdd) error {
	// 检查appId对应的应用是否存在
	app := entity.App{}
	if tx := global.MDB.Where(entity.App{Id: data.AppId}).First(&app); tx.RowsAffected == 0 {
		return errors.New("the corresponding appId '" + strconv.Itoa(int(data.AppId)) + "' does not exist")
	}

	// 构建实体
	var server entity.Server
	err := copier.Copy(&server, &data)
	if err != nil {
		return err
	}

	// 执行插入
	err = global.MDB.Create(&server).Error
	if err != nil {
		global.Log.Error("fail to add server！" + err.Error())
		return err
	}

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addServer"], string(remarks))
	return err
}

// GetServerAndApp 获取所有app以及对应server
func (appService *appService) GetServerAndApp() (err error, data []returnee.AppAndServerGet) {
	// 获取所有app
	var appSlice []entity.App
	if tx := global.MDB.Find(&appSlice); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, data
		} else {
			return errors.New("error querying app list: " + tx.Error.Error()), data
		}
	}

	// 获取所有服务
	var serverSlice []entity.Server
	if tx := global.MDB.Find(&serverSlice); tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return errors.New("error querying server list: " + tx.Error.Error()), data
	}

	// 拼接结果
	var serverMap = map[uint][]returnee.ServerSlice{}
	for _, serverItem := range serverSlice {
		var serverTemp returnee.ServerSlice
		_ = copier.Copy(&serverTemp, &serverItem)
		serverTemp.CreateTime = serverItem.CreateTime.Format("2006.01.02 15:04:05")
		serverMap[serverItem.AppId] = append(serverMap[serverItem.AppId], serverTemp)
	}
	for _, appItem := range appSlice {
		var appTemp returnee.AppAndServerGet
		_ = copier.Copy(&appTemp, &appItem)
		appTemp.CreateTime = appItem.CreateTime.Format("2006.01.02 15:04:05")
		appTemp.ServerSlice = serverMap[appItem.Id]
		data = append(data, appTemp)
	}

	return err, data
}
