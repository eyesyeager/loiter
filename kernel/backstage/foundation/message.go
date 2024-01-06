package foundation

import (
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/utils"
)

/**
 * 消息服务
 * @author eyesYeager
 * @date 2024/01/02
 */

type messageFoundation struct {
	emailAddr     string // 邮箱配置——SMTP服务器的地址
	emailIdentity string // 邮箱配置——身份证明
	emailUsername string // 邮箱配置——用户名
	emailPassword string // 邮箱配置——密码
	emailHost     string // 邮箱配置——主机地址
}

var MessageFoundation = messageFoundation{
	config.Program.EmailConfig.Addr,
	config.Program.EmailConfig.Identity,
	config.Program.EmailConfig.Username,
	config.Program.EmailConfig.Password,
	config.Program.EmailConfig.Host,
}

// SendEmailWithText 发送邮件(Text格式)
func (messageFoundation *messageFoundation) SendEmailWithText(subject string, from string, to string, cc string, text string) error {
	err := utils.SendEmailWithText(subject, from, to, cc, text,
		messageFoundation.emailAddr, messageFoundation.emailIdentity, messageFoundation.emailUsername, messageFoundation.emailPassword, messageFoundation.emailHost)
	if err != nil {
		global.BackstageLogger.Error("'SendEmailWithText' method error! failed to send email, error:", err.Error())
	}
	return err
}

// SendEmailWithHTML 发送邮件(HTML格式)
func (messageFoundation *messageFoundation) SendEmailWithHTML(subject string, from string, to string, cc string, html string) error {
	err := utils.SendEmailWithHTML(subject, from, to, cc, html,
		messageFoundation.emailAddr, messageFoundation.emailIdentity, messageFoundation.emailUsername, messageFoundation.emailPassword, messageFoundation.emailHost)
	if err != nil {
		global.BackstageLogger.Error("'SendEmailWithHTML' method error! failed to send email, error:", err.Error())
	}
	return err
}
