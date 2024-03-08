package capability

import (
	"encoding/json"
	"fmt"
	"loiter/config"
	"loiter/constants"
	"loiter/global"
	"loiter/model/entity"
	"loiter/utils"
)

/**
 * 消息通知能力
 * @auth eyesYeager
 * @date 2024/2/20 14:17
 */

type noticeFoundation struct {
	emailFrom     string // 邮箱配置——发件人
	emailAddr     string // 邮箱配置——SMTP服务器的地址
	emailIdentity string // 邮箱配置——身份证明
	emailUsername string // 邮箱配置——用户名
	emailPassword string // 邮箱配置——密码
	emailHost     string // 邮箱配置——主机地址
}

var NoticeFoundation = noticeFoundation{
	emailFrom:     config.Program.Name,
	emailAddr:     config.Program.EmailConfig.Addr,
	emailIdentity: config.Program.EmailConfig.Identity,
	emailUsername: config.Program.EmailConfig.Username,
	emailPassword: config.Program.EmailConfig.Password,
	emailHost:     config.Program.EmailConfig.Host,
}

// SendSiteNotice 发送站内通知
func (noticeFoundation *noticeFoundation) SendSiteNotice(host string, title string, content string, remarks string) {
	noticeFoundation.persistentNotice(host, title, content, constants.Notice.Site, false, remarks)
}

// SendEmailWithText 发送邮件(Text格式)
func (noticeFoundation *noticeFoundation) SendEmailWithText(host string, subject string, to []string, text string) error {
	return noticeFoundation.SendEmailWithTextAndCC(host, subject, to, []string{}, text, false)
}

// SendEmailWithTextAndCC 发送邮件(Text格式，附CC)
func (noticeFoundation *noticeFoundation) SendEmailWithTextAndCC(host string, subject string, to []string, cc []string, text string, secret bool) error {
	err := utils.SendEmail(subject, to, cc, text, "",
		noticeFoundation.emailFrom, noticeFoundation.emailAddr, noticeFoundation.emailIdentity, noticeFoundation.emailUsername, noticeFoundation.emailPassword, noticeFoundation.emailHost)
	if err != nil {
		global.AppLogger.Error("'SendEmailWithText' method error, failed to send email!"+
			"subject:", subject,
			";to:", to,
			";cc:", cc,
			";error:", err.Error())
	}
	remarksMap := make(map[string]interface{})
	remarksMap["to"] = to
	remarksMap["cc"] = cc
	remarks, _ := json.Marshal(remarksMap)
	noticeFoundation.persistentNotice(host, subject, text, constants.Notice.Email, secret, string(remarks))
	return err
}

// SendEmailWithHTML 发送邮件(HTML格式)
func (noticeFoundation *noticeFoundation) SendEmailWithHTML(host string, subject string, to []string, html string) error {
	return noticeFoundation.SendEmailWithHTMLAndCC(host, subject, to, []string{}, html, false)
}

// SendEmailWithHTMLAndCC 发送邮件(HTML格式，附CC)
func (noticeFoundation *noticeFoundation) SendEmailWithHTMLAndCC(host string, subject string, to []string, cc []string, html string, secret bool) error {
	err := utils.SendEmail(subject, to, cc, "", html,
		noticeFoundation.emailFrom, noticeFoundation.emailAddr, noticeFoundation.emailIdentity, noticeFoundation.emailUsername, noticeFoundation.emailPassword, noticeFoundation.emailHost)
	if err != nil {
		global.AppLogger.Error("'SendEmailWithHTML' method error, failed to send email!",
			"subject:", subject,
			";to:", to,
			";cc:", cc,
			";error:", err.Error())
	}
	remarksMap := make(map[string]interface{})
	remarksMap["to"] = to
	remarksMap["cc"] = cc
	remarks, _ := json.Marshal(remarksMap)
	noticeFoundation.persistentNotice(host, subject, html, constants.Notice.Email, secret, string(remarks))
	return err
}

// persistentNotice 消息持久化
func (noticeFoundation *noticeFoundation) persistentNotice(host string, title string, content string, genre string, secret bool, remarks string) {
	checkApp := entity.App{Host: host}
	if host != constants.Talisman.WithoutApp {
		if err := global.MDB.Where(&checkApp).First(&checkApp).Error; err != nil {
			checkApp.Name = "ERROR"
			global.GatewayLogger.Error(fmt.Sprintf("the network is abnormal or the application with host %s does not exist.", host))
		}
	}
	var notice = entity.Notice{
		AppName: checkApp.Name,
		Title:   title,
		Content: content,
		Genre:   genre,
		Secret:  secret,
		Remarks: remarks,
	}
	if err := global.MDB.Create(&notice).Error; err != nil {
		marshal, _ := json.Marshal(notice)
		global.GatewayLogger.Error(fmt.Sprintf("persistence notification failed, entity:%s, error:%s", string(marshal), err.Error()))
	}
}
