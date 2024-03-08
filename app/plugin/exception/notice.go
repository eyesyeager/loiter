package exception

import (
	"fmt"
	"loiter/app/capability"
	"loiter/constants/template"
	"loiter/global"
	"net/http"
)

/**
 * 异常通知插件
 * @auth eyesYeager
 * @date 2024/2/20 14:59
 */

func NoticeException(w http.ResponseWriter, r *http.Request, host string, errInfo string) {
	var email string
	if err := global.MDB.Raw(`SELECT u.email FROM app a, user u WHERE a.host = ? AND a.owner_id = u.id LIMIT 1`, host).Scan(&email).Error; err != nil {
		global.GatewayLogger.Error(fmt.Sprintf("failed to obtain the email address of the person responsible for the application whose host is %s. Error info: %s. Original notification information: %s", host, err.Error(), errInfo))
		return
	}
	emailTemplate := template.GetCommonEmailTemplate("处理器异常通知", errInfo)
	_ = capability.NoticeFoundation.SendEmailWithHTML(host, emailTemplate.Subject, []string{email}, emailTemplate.Content)
}
