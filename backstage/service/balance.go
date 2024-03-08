package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/global"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"time"
)

/**
 * 负载均衡业务层
 * @auth eyesYeager
 * @date 2024/1/5 16:51
 */

type balancerService struct {
}

var BalancerService = balancerService{}

// AddAppBalancer 添加应用负载均衡策略
func (*balancerService) AddAppBalancer(uid uint, appId uint, balancer string) error {
	if err := global.MDB.Create(&entity.AppBalancer{
		AppId:      appId,
		Balancer:   balancer,
		OperatorId: uid,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	return nil
}

// UpdateAppBalancer 更新应用负载均衡策略
func (*balancerService) UpdateAppBalancer(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.UpdateAppBalancer) error {
	// 获取应用当前策略
	var checkAppBalancer entity.AppBalancer
	if err := global.MDB.Where(&entity.AppBalancer{AppId: data.AppId}).First(&checkAppBalancer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("id为%d的应用没有配置负载均衡策略！", data.AppId))
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 如果新策略与当前策略相同，则取消更新
	if checkAppBalancer.Balancer == data.Balancer {
		return errors.New("更新策略与当前策略相同！")
	}
	// 校验是否存在对应策略
	checkBalancer := entity.Balancer{Code: data.Balancer}
	if err := global.MDB.Where(&checkBalancer).First(&checkBalancer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("无效负载均衡策略:%s", data.Balancer))
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 更新应用策略
	if err := global.MDB.Model(&entity.AppBalancer{}).Where("app_id", data.AppId).Updates(entity.AppBalancer{
		Balancer:   data.Balancer,
		OperatorId: userClaims.Uid,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 记录操作日志
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateAppBalancer, checkAppBalancer.Balancer, data.Balancer))
	return nil
}

// DeleteAppBalancer 删除应用负载均衡策略
func (*balancerService) DeleteAppBalancer(appId uint) error {
	return global.MDB.Where(entity.AppBalancer{
		AppId: appId,
	}).Unscoped().Delete(&entity.AppBalancer{}).Error
}

// GetAllBalancer 获取所有负载均衡策略
func (*balancerService) GetAllBalancer() (err error, res []returnee.GetDictionary) {
	var balancerList []entity.Balancer
	if err = global.MDB.Find(&balancerList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for _, item := range balancerList {
		res = append(res, returnee.GetDictionary{
			Label: item.Name,
			Value: item.Code,
		})
	}
	return err, res
}

// GetBalancerByPage 分页获取应用负载均衡信息
func (*balancerService) GetBalancerByPage(data receiver.GetBalancerByPage) (err error, res returnee.GetBalancerByPage) {
	// 获取应用负载均衡明细
	var balancerPOList []po.GetBalancerByPage
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = global.MDB.Raw(`SELECT ab.id, ab.balancer BalancerCode, ab.updated_at, b.name BalancerName, a.name AppName, u.username Operator 
				FROM app_balancer ab, balancer b, app a, user u 
				WHERE ab.app_id = a.id AND ab.balancer = b.code AND ab.operator_id = u.id AND (? = '' OR a.name = ?) AND (? = '' OR ab.balancer = ?)
				ORDER BY ab.updated_at DESC
				LIMIT ?, ?`, data.AppName, data.AppName, data.Balancer, data.Balancer, offset, limit).Scan(&balancerPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	// 时间处理
	var balancerList []returnee.GetBalancerByPageInner
	for _, item := range balancerPOList {
		var t returnee.GetBalancerByPageInner
		_ = copier.Copy(&t, &item)
		t.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
		balancerList = append(balancerList, t)
	}
	// 查询总数
	var total int64
	if len(balancerPOList) < data.PageSize {
		total = int64(len(balancerPOList))
	} else {
		if err = global.MDB.Model(&entity.AppBalancer{}).Where(&entity.AppBalancer{}).Count(&total).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
		}
		if err = global.MDB.Raw(`SELECT COUNT(*) FROM app_balancer ab, app a 
                WHERE ab.app_id = a.id AND (? = '' OR a.name = ?) AND (? = '' OR ab.balancer = ?)`,
			data.AppName, data.AppName, data.Balancer, data.Balancer).Scan(&total).Error; err != nil {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
		}
	}
	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = balancerList
	return err, res
}
