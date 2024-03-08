package template

import (
	"loiter/config"
	"loiter/model/structure"
	"time"
)

/**
 * 通用邮件模版
 * @auth eyesYeager
 * @date 2024/2/23 11:02
 */

// GetCommonEmailTemplate 通用邮件模版
func GetCommonEmailTemplate(subject string, content string) structure.EmailReturnStruct {
	systemName := config.Program.Name
	nowTime := time.Now().Format(time.DateTime)
	return structure.EmailReturnStruct{
		Subject: subject,
		Content: `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>` + systemName + ` | ` + subject + `</title>
  </head>

  <body style="margin: 0; padding: 0;">
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="border-bottom: 2px solid rgb(73, 69, 0);">
            <td style="font-size: 25px;color: rgb(73, 69, 0);">` + systemName + ` | ` + subject + `</td>
          </tr>
        </table>
      </tr>
      
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="height: 20px;"></tr>
          <tr>
            <td>` + content + `</td>
          </tr>
        </table>
      </tr>

      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="height: 150px;"></tr>
          <tr>
            <td align="right" style="height: 30px;">` + systemName + `</td>
          </tr>
          <tr>
            <td align="right">` + nowTime + `</td>
          </tr>
        </table>
      </tr>
    </table>
  </body>
</html>
`,
	}
}
