package service

import (
	"loiter/global"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/backstage/utils"
	"net/http"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/9/27 12:36
 */

type logService struct {
}

var LogService = logService{}

// Operate 操作日志
func (*logService) Operate() {

}

// Login 登录日志
func (*logService) Login(r *http.Request, uid uint) {
	logLogin := entity.LogLogin{
		Uid:     uid,
		Ip:      utils.GetIp(r),
		Browser: utils.GetBrowser(r),
	}

	// 插入数据库
	if err := global.MDB.Create(&logLogin).Error; err != nil {
		global.BackstageLogger.Error("User login log insertion failed for userId " + strconv.Itoa(int(uid)) + ", error:" + err.Error())
	}
}
