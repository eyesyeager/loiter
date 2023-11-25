package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"loiter/config"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/service"
	"loiter/kernel/utils"
	"net/http"
)

/**
 * 用户模块控制器
 * @author eyesYeager
 * @date 2023/9/26 14:44
 */

// DoRegister 用户注册
func DoRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hash, err := bcrypt.GenerateFromPassword([]byte("loiter"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash)
	fmt.Println(encodePWD)
}

// DoLogin 用户登录
func DoLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 参数校验
	var data receiver.DoLogin
	if err := parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 还原原始密码
	if err, decrypt := utils.AesDecrypt(data.Password, config.Program.AESSecretKey); err == nil {
		data.Password = decrypt
	} else {
		result.FailAttachedMsg(w, r, "illegal password encryption method,"+err.Error())
	}

	// 执行业务
	if err := service.UserService.DoLogin(w, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
