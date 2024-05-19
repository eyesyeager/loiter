package filter

import (
	"errors"
	"fmt"
	"loiter/app/capability"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/container"
	"loiter/utils"
	"net/http"
)

/**
 * 黑白名单
 * @auth eyesYeager
 * @date 2024/1/11 16:47
 */

func NameListFilter(w http.ResponseWriter, r *http.Request, host string, genre string) (error, bool) {
	genres := container.NameListByAppMap[host]
	ip := utils.GetIp(r)
	for _, item := range genres {
		var err error
		var success bool
		// 黑白名单校验
		if item == constants.NameList.White.Value { // 白名单校验
			err, success = checkWhiteNameList(host, ip)
		} else if item == constants.NameList.Black.Value { // 黑名单校验
			err, success = checkBlackNameList(host, ip)
		} else {
			return errors.New(fmt.Sprintf("there is no nameList of type %s", item)), false
		}
		// 处理返回值
		if err != nil {
			return err, success
		}
		if !success {
			errMsg := fmt.Sprintf("application access with host %s is blocked by the nameList, genre: %s, ip: %s",
				host, item, ip)
			statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.Forbidden, errMsg, genre)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			go capability.NoticeFoundation.SendSiteNotice(host, "黑白名单触发拦截", errMsg,
				fmt.Sprintf("请求路径：%s；", r.URL.Path))
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
