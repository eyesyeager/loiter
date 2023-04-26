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
 * @author eyesYeager
 * @date 2023/4/26 16:56
 */

type serverService struct {
}

var ServerService = serverService{}

// AddServer 添加Server
func (serverService *serverService) AddServer(r *http.Request, data receiver.ServerAdd) error {
	// 检查BasketId对应的组是否存在
	if tx := global.MDB.Where(entity.App{Id: data.BasketId}).First(&entity.Basket{}); tx.RowsAffected == 0 {
		return errors.New("the corresponding basketId '" + strconv.Itoa(int(data.BasketId)) + "' does not exist")
	}

	// 检查是否存在同地址Server
	checkServer := entity.Server{}
	if tx := global.MDB.Where(entity.Server{
		BasketId: data.BasketId,
		Server:   data.Server,
	}).First(&checkServer); tx.RowsAffected != 0 {
		return errors.New("the server '" + data.Server + "' has already been registered with ID " + strconv.Itoa(int(checkServer.Id)))
	}

	// 构建实体
	var server entity.Server
	if err := copier.Copy(&server, &data); err != nil {
		return err
	}

	// 执行插入
	if err := global.MDB.Create(&server).Error; err != nil {
		global.Log.Error("fail to add server！" + err.Error())
		return err
	}

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addServer"], string(remarks))
	return nil
}
