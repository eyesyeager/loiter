package email

import (
	"loiter/kernel/backstage/model/structure"
)

// GetRegisterEmailTemplate 账号开通邮件模版
func GetRegisterEmailTemplate() structure.EmailReturnStruct {
	return structure.EmailReturnStruct{
		Subject: "账号开通通知",
		Content: `账号开通？？？`,
	}
}
