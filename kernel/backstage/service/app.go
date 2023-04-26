package service

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
	"zliway/global"
	"zliway/kernel/backstage/model/entity"
	"zliway/kernel/backstage/model/receiver"
)

/**
 * app相关业务
 * @author eyesYeager
 * @date 2023/4/11 17:57
 */

type appService struct {
}

var AppService = appService{}

// AddApp 添加应用
func (appService *appService) AddApp(r *http.Request, data receiver.AppAdd) error {
	// 检查应用是否已经注册
	var checkApp entity.App
	if tx := global.MDB.Where(entity.App{App: data.App}).Find(&checkApp); tx.RowsAffected != 0 {
		return errors.New("the application '" + data.App + "' has already been registered with ID " + strconv.Itoa(int(checkApp.Id)))
	}

	// 构建实体
	var app entity.App
	if err := copier.Copy(&app, &data); err != nil {
		return err
	}

	// 执行插入
	if err := global.MDB.Create(&app).Error; err != nil {
		global.Log.Error("fail to add app！" + err.Error())
		return err
	}

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addApp"], string(remarks))
	return nil
}
