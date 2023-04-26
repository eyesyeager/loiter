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
 * basket相关业务
 * @author eyesYeager
 * @date 2023/4/26 16:56
 */

type basketService struct {
}

var BasketService = basketService{}

// AddBasket 添加Basket
func (basketService *basketService) AddBasket(r *http.Request, data receiver.BasketAdd) error {
	// 检查appId对应的应用是否存在
	if tx := global.MDB.Where(entity.App{Id: data.AppId}).First(&entity.App{}); tx.RowsAffected == 0 {
		return errors.New("the corresponding appId '" + strconv.Itoa(int(data.AppId)) + "' does not exist")
	}

	// 检查是否存在同名Basket
	checkBasket := entity.Basket{}
	if tx := global.MDB.Where(entity.Basket{
		AppId: data.AppId,
		Name:  data.Name,
	}).First(&checkBasket); tx.RowsAffected != 0 {
		return errors.New("the basket '" + data.Name + "' has already been registered with ID " + strconv.Itoa(int(checkBasket.Id)))
	}

	// 构建实体
	var basket entity.Basket
	if err := copier.Copy(&basket, &data); err != nil {
		return err
	}

	// 执行插入
	if err := global.MDB.Create(&basket).Error; err != nil {
		global.Log.Error("fail to add basket！" + err.Error())
		return err
	}

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addBasket"], string(remarks))
	return nil
}
