package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/service"
	"net/http"
)

/**
 * 用户模块控制器
 * @author eyesYeager
 * @date 2023/9/26 14:44
 */

// DoRegister
// @Summary			用户注册
// @Description		admin 可创建 user，superAdmin 可创建 admin 和 user
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token				header		string		true		"身份令牌"
// @Param			receiver.DoRegister	body		string		false		"请求参数"
// @Success			200					{object}	result.Response
// @Failure			400					{object}	result.Response
// @Router			/user/doRegister [post]
func DoRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hash, err := bcrypt.GenerateFromPassword([]byte("loiter"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash)
	fmt.Println(encodePWD)
}

// DoLogin
// @Summary			用户登录
// @Description		权限：user
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token				header		string		true		"身份令牌"
// @Param			receiver.DoLogin	body		string		false		"请求参数"
// @Success			200					{object}	result.Response
// @Failure			400					{object}	result.Response
// @Router			/user/doLogin [post]
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

	// 执行业务
	if err := service.UserService.DoLogin(w, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
