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
	"zliway/kernel/register/constant"
)

/**
 * predicates相关业务
 * @author eyesYeager
 * @date 2023/4/26 23:05
 */

type predicatesService struct {
}

var PredicatesService = predicatesService{}

// AddPredicates 添加predicates
func (predicatesService *predicatesService) AddPredicates(r *http.Request, data receiver.PredicatesAdd) error {
	// 检查BasketId对应的组是否存在
	if tx := global.MDB.Where(entity.App{Id: data.BasketId}).First(&entity.Basket{}); tx.RowsAffected == 0 {
		return errors.New("the corresponding basketId '" + strconv.Itoa(int(data.BasketId)) + "' does not exist")
	}

	// 检查app是否是微服务模式
	checkApp := entity.App{}
	if err := global.MDB.Raw("select app.id, app.pattern from app, basket where basket.id=? and basket.app_id=app.id", data.BasketId).Scan(&checkApp).Error; err != nil {
		global.Log.Error("fail to get app: " + err.Error())
		return errors.New(err.Error())
	}
	if checkApp.Pattern != constant.AppPattern["micro"] {
		return errors.New("predicates can only be configured when the app is in microservice mode, The current app's ID is '" + strconv.Itoa(int(checkApp.Id)) + "', and the mode is " + strconv.Itoa(int(checkApp.Pattern)))
	}

	// 检查是否存在相同path
	var checkPredicatesId uint
	if tx := global.MDB.Raw("select predicates.id from app, basket, predicates where app.id=? and basket.app_id=app.id and predicates.basket_id=basket.id and predicates.path=?", checkApp.Id, data.Path).Scan(&checkPredicatesId); tx.RowsAffected != 0 {
		return errors.New("the path '" + data.Path + "' has already been registered with ID " + strconv.Itoa(int(checkPredicatesId)))
	}

	// 构建实体
	var predicates entity.Predicates
	if err := copier.Copy(&predicates, &data); err != nil {
		return err
	}

	// 执行插入
	if err := global.MDB.Create(&predicates).Error; err != nil {
		global.Log.Error("fail to add predicates！" + err.Error())
		return err
	}

	// 添加操作日志
	remarks, _ := json.Marshal(data)
	go LogService.AppOperateLog(r, LogService.OperateType["addPredicates"], string(remarks))
	return nil
}
