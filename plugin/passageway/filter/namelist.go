package filter

import (
	"errors"
	"fmt"
	"loiter/constant"
	"loiter/global"
	"loiter/helper"
	"loiter/kernel/container"
	"loiter/plugin/passageway/filter/namelist"
	"net/http"
)

/**
 * 黑白名单
 * @auth eyesYeager
 * @date 2024/1/11 16:47
 */

func NameListFilter(w http.ResponseWriter, r *http.Request, host string) (error, bool) {
	genres := container.NameListByAppMap[host]
	ip := helper.GetIp(r)
	for _, item := range genres {
		var err error
		var success bool
		// 黑白名单校验
		if item == namelist.WhiteList { // 白名单校验
			err, success = checkWhiteNameList(host, ip)
		} else if item == namelist.BlackList { // 黑名单校验
			err, success = checkBlackNameList(host, ip)
		} else {
			return errors.New(fmt.Sprintf("there is no nameList of type %s", item)), false
		}
		// 处理返回值
		if err != nil {
			return err, success
		}
		if !success {
			statusCode, contentType, content := helper.HtmlSimpleTemplate(constant.ResponseTitle.Forbidden, constant.ResponseNotice.Empty)
			helper.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(fmt.Sprintf("application access with host %s is blocked by the nameList, genre: %s, ip: %s", host, item, ip))
			return nil, false
		}
	}
	return nil, true
}

// checkWhiteNameList 校验白名单
func checkWhiteNameList(host string, ip string) (error, bool) {
	iNameList := container.WhiteNameListByAppMap[host]
	err, allow := iNameList.Check(ip)
	if err != nil {
		return errors.New(fmt.Sprintf("whitelist verification error! host: %s, ip: %s, error: %s", host, ip, err.Error())), false
	}
	return nil, allow
}

// checkBlackNameList 校验黑名单
func checkBlackNameList(host string, ip string) (error, bool) {
	iNameList := container.BlackNameListByAppMap[host]
	err, allow := iNameList.Check(ip)
	if err != nil {
		return errors.New(fmt.Sprintf("blacklist verification error! host: %s, ip: %s, error: %s", host, ip, err.Error())), false
	}
	return nil, allow
}
