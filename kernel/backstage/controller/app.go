package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/service"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 应用控制器
 * @author eyesYeager
 * @date 2023/4/11 17:55
 */

// AddApp
// @Summary			注册应用
// @Description		权限：admin
// @Tags			app
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token					header		string		true		"身份令牌"
// @Param			receiver.AddApp			body		string		false		"请求参数"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/app/addApp [post]
func AddApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.AddApp
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.AppService.AddApp(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}

// TODO：删除应用
