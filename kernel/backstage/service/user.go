package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/model/entity"
	"loiter/kernel/backstage/model/po"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/template/email"
	"loiter/kernel/backstage/utils"
	"net/http"
)

/**
 * @author eyesYeager
 * @date 2023/9/26 15:33
 */

type userService struct {
}

var UserService = userService{}

// DoLogin 用户登录
func (*userService) DoLogin(w http.ResponseWriter, r *http.Request, data receiver.DoLogin) error {
	errorMsg := "账号或密码错误"

	// 密码应该避免明文传输，因此前端使用了 AES 双向加密算法对密码加密
	// 因此应该先进行 AES 解密得到原始密码
	if err, decrypt := utils.AesDecrypt(data.Password, config.Program.AESSecretKey); err == nil {
		data.Password = decrypt
	} else {
		global.BackstageLogger.Warn("password ", data.Password, " does not comply with decryption rules, error:", err.Error())
		return errors.New(errorMsg)
	}

	// 获取用户密码
	var checkUser po.LoginUserRole
	if tx := global.MDB.Raw(
		"SELECT user.id Uid, user.password, role.name Role FROM user, role WHERE user.username = ? AND user.rid = role.id",
		data.Username).Scan(&checkUser); tx.RowsAffected != 1 {
		global.BackstageLogger.Warn("a user with username ", data.Username, " does not exist")
		return errors.New(errorMsg)
	}

	// 密码校验
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(data.Password))
	if err != nil {
		global.BackstageLogger.Warn("the password with username ", data.Username, " is not the given value")
		return errors.New(errorMsg)
	}

	// 生成token
	if err = foundation.AuthFoundation.RefreshToken(w, checkUser.Uid, checkUser.Role); err != nil {
		global.BackstageLogger.Warn("token generation failed for user with username ", data.Username, ", error:"+err.Error())
		return errors.New("令牌生成失败，请联系管理员处理")
	}

	// 添加登录日志
	go LogService.Login(r, checkUser.Uid)
	return err
}

// DoRegister 开通新账号
func (*userService) DoRegister(r *http.Request, userClaims utils.JwtCustomClaims, data receiver.DoRegister) (err error) {
	// TODO：验证码校验

	// 校验操作可行性，高级别用户只可创建低级别用户
	var compareResult int
	if err, compareResult = foundation.RoleFoundation.CompareRole(userClaims.Role, data.Role); err != nil {
		global.BackstageLogger.Warn("permission judgment error, incorrect data present, error:", err.Error())
		return errors.New("角色非法，系统不存在类型为 " + data.Role + " 的角色")
	}
	if compareResult <= 0 {
		return errors.New("您的权限不足以创建类型为 " + data.Role + " 的角色")
	}

	// 生成随机密码
	initialPsd := utils.GenerateRandString(config.Context.InitialPsdLen)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(initialPsd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	// 创建用户
	_, rid := foundation.RoleFoundation.GetRidByRole(data.Role)
	user := entity.User{
		Username: data.Username,
		Password: string(passwordHash),
		Rid:      rid,
		Email:    data.Email,
		Remarks:  data.Remarks,
	}
	if err = global.MDB.Create(&user).Error; err != nil {
		global.BackstageLogger.Warn("failed to create user, error:", err.Error())
		return errors.New("创建用户失败，请联系管理员")
	}

	// 发送邮件通知被创建的用户
	// 此处不能异步，因为邮件中含有初始密码，因此必须保证邮件发送成功
	template := email.GetRegisterEmailTemplate(config.Program.Name, data.Username, initialPsd)
	err = foundation.MessageFoundation.SendEmailWithHTML(template.Subject, []string{data.Email}, template.Content)
	if err != nil {
		global.BackstageLogger.Error("email sending failed when registering a new user,",
			"username:", data.Username,
			";email:", data.Email,
			";error:", err.Error())
		return errors.New("邮件通知新用户失败！请记住初始密码：" + initialPsd)
	}

	// 记录操作日志
	go LogService.Universal(r, userClaims.Uid,
		constant.BuildUniversalLog(constant.LogUniversal.DoRegister, user.Username, user.Email, user.Remarks))
	return err
}
