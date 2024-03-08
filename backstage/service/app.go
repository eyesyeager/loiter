package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"loiter/backstage/constant"
	"loiter/backstage/controller/result"
	"loiter/config"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/model/entity"
	"loiter/model/po"
	"loiter/model/receiver"
	"loiter/model/returnee"
	"loiter/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @author eyesYeager
 * @date 2023/9/26 15:33
 */

type appService struct {
}

var AppService = appService{}

// SaveApp 注册/编辑应用
func (a *appService) SaveApp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.SaveApp) error {
	// 参数校验
	serverList := data.ServerList
	if serverList == nil || len(serverList) == 0 {
		return errors.New("服务实例不能为空")
	}
	// 执行保存操作
	if data.AppId == 0 { // 新增
		if err := a.addApp(r, userClaims, data); err != nil {
			return err
		}
	} else { // 编辑
		if err := a.updateApp(r, userClaims, data); err != nil {
			return err
		}
	}
	return nil
}

// addApp 注册app
func (*appService) addApp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.SaveApp) error {
	// 检查Host的唯一性
	checkApp := entity.App{}
	if tx := global.MDB.Where(&entity.App{Host: data.Host}).First(&checkApp); tx.RowsAffected != 0 {
		return errors.New(fmt.Sprintf("host为'%s'的应用已存在，应用名为'%s'", data.Host, checkApp.Name))
	}
	// 插入应用
	var newApp = entity.App{
		Name:    data.AppName,
		Host:    data.Host,
		OwnerId: userClaims.Uid,
		Remarks: data.Remarks,
	}
	if err := global.MDB.Create(&newApp).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 插入服务实例
	var newServerList []entity.Server
	serverList := data.ServerList
	for _, item := range serverList {
		newServerList = append(newServerList, entity.Server{
			AppId:   newApp.ID,
			Address: item.Address,
			Weight:  item.Weight,
			Status:  constant.Status.Normal.Code,
		})
	}
	if err := global.MDB.Create(&newServerList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 插入默认负载均衡策略
	if err := BalancerService.AddAppBalancer(userClaims.Uid, newApp.ID, config.Program.PluginConfig.BalancerDefaultStrategy); err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	remark := ""
	// 插入默认过滤器
	if config.Program.PluginConfig.FilterDefaultStrategy != "" {
		if err := global.MDB.Create(&entity.AppProcessor{
			AppId: newApp.ID,
			Genre: constants.Processor.Filter.Code,
			Codes: config.Program.PluginConfig.FilterDefaultStrategy,
		}).Error; err != nil {
			remark += err.Error() + ","
		}
	}
	// 插入默认响应处理器
	if config.Program.PluginConfig.AidDefaultStrategy != "" {
		if err := global.MDB.Create(&entity.AppProcessor{
			AppId: newApp.ID,
			Genre: constants.Processor.Aid.Code,
			Codes: config.Program.PluginConfig.AidDefaultStrategy,
		}).Error; err != nil {
			remark += err.Error() + ","
		}
	}
	// 插入默认异常处理器
	if config.Program.PluginConfig.ExceptionDefaultStrategy != "" {
		if err := global.MDB.Create(&entity.AppProcessor{
			AppId: newApp.ID,
			Genre: constants.Processor.Exception.Code,
			Codes: config.Program.PluginConfig.ExceptionDefaultStrategy,
		}).Error; err != nil {
			remark += err.Error() + ","
		}
	}
	// 插入最终处理器
	if config.Program.PluginConfig.FinalDefaultStrategy != "" {
		if err := global.MDB.Create(&entity.AppProcessor{
			AppId: newApp.ID,
			Genre: constants.Processor.Final.Code,
			Codes: config.Program.PluginConfig.FinalDefaultStrategy,
		}).Error; err != nil {
			remark += err.Error() + ","
		}
	}
	// 记录操作日志
	marshal, _ := json.Marshal(data)
	go LogService.App(r, userClaims.Uid, newApp.ID,
		constant.BuildUniversalLog(constant.LogUniversal.AddApp, marshal, remark))
	return nil
}

// updateApp 编辑app
func (*appService) updateApp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.SaveApp) error {
	// 检查Host的唯一性
	checkApp := entity.App{}
	if tx := global.MDB.Where("id != ? AND host = ?", data.AppId, data.Host).First(&checkApp); tx.RowsAffected != 0 {
		return errors.New(fmt.Sprintf("host为'%s'的应用已存在，应用名为'%s'", data.Host, checkApp.Name))
	}
	// 编辑app(如果Model可以作为条件，记得改其他地方)
	if err := global.MDB.Model(&entity.App{Model: gorm.Model{ID: data.AppId}}).Updates(entity.App{
		Model:   gorm.Model{ID: data.AppId},
		Name:    data.AppName,
		Host:    data.Host,
		OwnerId: data.OwnerId,
		Remarks: data.Remarks,
	}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 重建server
	if err := global.MDB.Where(&entity.Server{AppId: data.AppId}).Unscoped().Delete(&entity.Server{}).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	var newServerList []entity.Server
	serverList := data.ServerList
	for _, item := range serverList {
		newServerList = append(newServerList, entity.Server{
			AppId:   data.AppId,
			Address: item.Address,
			Weight:  item.Weight,
			Status:  constant.Status.Normal.Code,
		})
	}
	if err := global.MDB.Create(&newServerList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	marshal, _ := json.Marshal(data)
	go LogService.App(r, userClaims.Uid, data.AppId,
		constant.BuildUniversalLog(constant.LogUniversal.UpdateApp, marshal))
	return nil
}

// ActivateApp 激活/失效应用
func (*appService) ActivateApp(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.ActivateApp) error {
	// 应用存在性校验
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, data.AppId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("不存在id为%d的应用", data.AppId))
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 更新应用状态
	var targetStatus uint8
	if checkApp.Status == constant.Status.Normal.Code {
		targetStatus = constant.Status.Invalid.Code
	} else if checkApp.Status == constant.Status.Invalid.Code {
		targetStatus = constant.Status.Normal.Code
	} else {
		return errors.New(fmt.Sprintf("非法应用状态：%s", constant.Status.GetNameByCode(checkApp.Status)))
	}
	if err := global.MDB.Model(&checkApp).Where(&checkApp).Update("status", targetStatus).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid, constant.BuildUniversalLog(
		constant.LogUniversal.ActivateApp, checkApp.Name,
		constant.Status.GetNameByCode(checkApp.Status),
		constant.Status.GetNameByCode(targetStatus),
	))
	return nil
}

// DeleteApp 删除应用
func (*appService) DeleteApp(r *http.Request, userClaims utils.JwtCustomClaims, appId uint) error {
	// 校验appId
	var checkApp entity.App
	if err := global.MDB.First(&checkApp, appId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("不存在id为%d的应用", appId))
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
		}
	}
	// 删除应用
	if err := global.MDB.Unscoped().Delete(&entity.App{}, appId).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error()))
	}
	remarks := ""
	// 删除应用负载均衡器
	if err := BalancerService.DeleteAppBalancer(appId); err != nil {
		remarks += fmt.Sprintf("删除负载均衡器失败，错误信息:%s;", err.Error())
	}
	// 删除应用处理器
	if err := ProcessorService.DeleteAppProcessor(appId); err != nil {
		remarks += fmt.Sprintf("删除处理器失败，错误信息:%s;", err.Error())
	}
	// 删除容器
	container.DeleteRegister(checkApp.Host)
	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DeleteApp, checkApp.Name, remarks))
	return nil
}

// GetAllApp 获取所有应用信息
func (*appService) GetAllApp() (error, []string) {
	var appList []entity.App
	if err := global.MDB.Find(&appList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), nil
	}
	var appNameList []string
	for _, item := range appList {
		appNameList = append(appNameList, item.Name)
	}
	// 如果没有应用，就返回空数组，而不是nil
	if appNameList == nil {
		appNameList = []string{}
	}
	return nil, appNameList
}

// GetAppInfoByPage 分页获取应用信息
func (*appService) GetAppInfoByPage(data receiver.GetAppInfoByPage) (err error, res returnee.GetAppInfoByPage) {
	var status uint8
	if data.Status == "" {
		status = 0
	} else {
		num, err := strconv.ParseUint(data.Status, 10, 8)
		if err != nil {
			return errors.New(fmt.Sprintf("status(%s)参数格式错误！", data.Status)), res
		} else {
			status = uint8(num)
		}
	}
	// 查询明细主体
	var resPOList []po.GetAppInfoByPage
	limit, offset := utils.BuildPageSearch(data.PageStruct)
	if err = global.MDB.Raw(`SELECT a.id AppId, a.name AppName, a.host, a.status, a.remarks, a.created_at, ab.balancer, u.username Owner
							FROM app a
							LEFT JOIN app_balancer ab ON a.id = ab.app_id 
							LEFT JOIN user u ON a.owner_id = u.id
							WHERE (? = '' OR a.name = ?) AND (? = 0 OR a.status = ?) 
							ORDER BY a.created_at DESC
							LIMIT ?, ?`, data.AppName, data.AppName, status, status, offset, limit).Scan(&resPOList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	if len(resPOList) == 0 {
		return nil, res
	}
	// 获取服务实例数据
	var serverList []entity.Server
	if err = global.MDB.Find(&serverList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	var validNumByServer = make(map[uint]int)
	var numByServer = make(map[uint]int)
	for _, item := range serverList {
		if item.Status == constant.Status.Normal.Code {
			validNumByServer[item.AppId] += 1
		}
		numByServer[item.AppId] += 1
	}
	// 获取应用处理器数据
	var numByProcessor = make(map[uint]int)
	var processors []entity.AppProcessor
	if err = global.MDB.Find(&processors).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	for _, item := range processors {
		split := strings.Split(item.Codes, config.Program.PluginConfig.ProcessorDelimiter)
		numByProcessor[item.AppId] += len(split)
	}

	// 构建负载均衡器翻译map
	err, balancerDict := BalancerService.GetAllBalancer()
	if err != nil {
		return err, res
	}
	balancerMap := make(map[string]string)
	for _, item := range balancerDict {
		balancerMap[item.Value] = item.Label
	}

	// 明细主体构建
	var innerResList []returnee.GetAppInfoByPageInner
	for _, item := range resPOList {
		var innerRes returnee.GetAppInfoByPageInner
		_ = copier.Copy(&innerRes, &item)
		innerRes.Status = constant.Status.GetNameByCode(item.Status)
		innerRes.CreatedAt = item.CreatedAt.Format(time.DateOnly)
		innerRes.ServerNum = numByServer[item.AppId]
		innerRes.ValidServerNum = validNumByServer[item.AppId]
		innerRes.Balancer = balancerMap[innerRes.Balancer]
		innerRes.Plugins = numByProcessor[item.AppId]
		innerResList = append(innerResList, innerRes)
	}
	// 查询总数
	var total int64
	if err = global.MDB.Model(&entity.App{}).Where(&entity.App{
		Name:   data.AppName,
		Status: status,
	}).Count(&total).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	res.PageStruct = data.PageStruct
	res.Total = total
	res.Data = innerResList
	return err, res
}

// GetAppInfoById 根据id获取应用信息
func (*appService) GetAppInfoById(appId uint) (err error, res returnee.GetAppInfoById) {
	// 获取应用基本信息
	var checkApp entity.App
	if err = global.MDB.First(&checkApp, appId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("不存在id为%d的应用", appId)), res
		} else {
			return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
		}
	}
	// 获取应用服务实例信息
	var serverEntityList []entity.Server
	if err = global.MDB.Where(&entity.Server{AppId: appId}).Find(&serverEntityList).Error; err != nil {
		return errors.New(fmt.Sprintf(result.CommonInfo.DbOperateError, err.Error())), res
	}
	var serverList []returnee.AppServer
	for _, item := range serverEntityList {
		serverList = append(serverList, returnee.AppServer{
			Address: item.Address,
			Weight:  item.Weight,
		})
	}
	return err, returnee.GetAppInfoById{
		AppId:      appId,
		AppName:    checkApp.Name,
		Host:       checkApp.Host,
		OwnerId:    checkApp.OwnerId,
		ServerList: serverList,
		Remarks:    checkApp.Remarks,
	}
}
