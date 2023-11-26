package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"loiter/config"
	"loiter/global"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/model/po"
	"loiter/kernel/backstage/model/receiver"
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
		global.BackstageLogger.Warn("password " + data.Password + " does not comply with decryption rules, error:" + err.Error())
		return errors.New(errorMsg)
	}

	// 获取用户密码
	var checkUser po.LoginUserRole
	if tx := global.MDB.Raw(
		"SELECT user.id UserId, user.password, role.weight FROM user, role WHERE user.username = ? AND user.role_id = role.id",
		data.Username).Scan(&checkUser); tx.RowsAffected != 1 {
		global.BackstageLogger.Warn("a user with username " + data.Username + " does not exist")
		return errors.New(errorMsg)
	}

	// 密码校验
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(data.Password))
	if err != nil {
		global.BackstageLogger.Warn("the password with username " + data.Username + " is not " + data.Password)
		return errors.New(errorMsg)
	}

	// 生成token
	var token string
	if token, err = utils.GenerateToken(checkUser.UserId, checkUser.Weight); err != nil {
		global.BackstageLogger.Warn("token generation failed for user with username " + data.Username + "; error info:" + err.Error())
		return errors.New("令牌生成失败:" + err.Error())
	} else {
		w.Header().Set(constant.ResponseHead.Token, token)
	}

	// 添加登录日志
	go LogService.Login(r, checkUser.UserId, token)
	return nil
}

// DoRegister 用户注册
func (*userService) DoRegister(w http.ResponseWriter, r *http.Request, data receiver.DoRegister) error {
	// 验证码校验

	//
	hash, err := bcrypt.GenerateFromPassword([]byte("loiter"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash)
	fmt.Println(encodePWD)
	return nil
}
