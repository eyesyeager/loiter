package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/kernel/backstage/constant"
	"loiter/kernel/backstage/controller/parser"
	"loiter/kernel/backstage/controller/result"
	"loiter/kernel/backstage/controller/validator"
	"loiter/kernel/backstage/foundation"
	"loiter/kernel/backstage/model/receiver"
	"loiter/kernel/backstage/service"
	"net/http"
)

/**
 * 负载均衡控制器
 * @auth eyesYeager
 * @date 2024/1/5 16:43
 */

// UpdateAppBalance
// @Summary			更新应用负载均衡策略
// @Description		权限：admin
// @Tags			log
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.UpdateAppBalance			body		string		false		"请求参数"
// @Success			200						{object}	result.Response
// @Failure			400						{object}	result.Response
// @Router			/balance/updateAppBalance [post]
func UpdateAppBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppBalance
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.BalanceService.UpdateAppBalance(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
