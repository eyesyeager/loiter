package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/model/receiver"
	"net/http"
)

/**
 * 用户模块控制器
 * @author eyesYeager
 * @date 2023/9/26 14:44
 */

// DoLogin
// @Summary			用户登录
// @Description		权限：visitor
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
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err := service.UserService.DoLogin(w, r, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// DoRegister
// @Summary			开通新账号
// @Description		操作人可创建比自己权限等级低的用户
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
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		return
	}

	// 参数校验
	var data receiver.DoRegister
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, err.Error())
		return
	}

	// 执行业务
	if err = service.UserService.DoRegister(r, userClaims, data); err == nil {
		result.SuccessDefault(w, nil)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetUserInfo
// @Summary			获取用户信息
// @Description		权限：user
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token				header		string		true		"身份令牌"
// @Success			200					{object}	result.Response
// @Failure			400					{object}	result.Response
// @Router			/user/getUserInfo [get]
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User)
	if err != nil {
		return
	}

	// 执行业务
	if err, res := service.UserService.GetUserInfo(userClaims); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}

// GetAllUser
// @Summary			获取所有用户
// @Description		权限：user
// @Tags			user
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token				header		string		true		"身份令牌"
// @Success			200					{object}	result.Response
// @Failure			400					{object}	result.Response
// @Router			/user/getAllUser [get]
func GetAllUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if _, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.User); err != nil {
		return
	}

	// 执行业务
	if err, res := service.UserService.GetAllUser(); err == nil {
		result.SuccessDefault(w, res)
	} else {
		result.FailAttachedMsg(w, err.Error())
	}
}
