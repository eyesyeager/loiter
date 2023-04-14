package service

import (
	"encoding/json"
	"net/http"
	"zliway/global"
	"zliway/kernel/backstage/model/entity"
	"zliway/kernel/utils"
)

/**
 * 操作日志业务
 * @author eyesYeager
 * @date 2023/4/14 13:57
 */

type logService struct {
	OperateType map[string]string
}

// AppOperateLog app相关操作日志
func (logService *logService) AppOperateLog(r *http.Request, operateType string, remarks string) {
	log := entity.LogOperate{
		Pattern: operateType,
		Remarks: remarks,
		Ip:      utils.GetIp(r),
		Browser: utils.GetBrowser(r),
	}
	if err := global.MDB.Create(&log).Error; err != nil {
		logStr, _ := json.Marshal(log)
		global.Log.Error("operation log insertion failed, err: " + err.Error() + " log:" + string(logStr))
	}
}

// LogService 对外暴露实例
var LogService = logService{
	OperateType: map[string]string{
		"addApp":           "添加应用",
		"updateApp":        "修改应用",
		"deleteApp":        "删除应用",
		"onlineApp":        "上线应用",
		"deactivateApp":    "停用应用",
		"addServer":        "添加服务",
		"updateServer":     "修改服务",
		"deleteServer":     "删除服务",
		"onlineServer":     "上线服务",
		"deactivateServer": "停用服务",
	},
}
