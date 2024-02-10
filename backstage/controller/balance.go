package controller

import (
	"github.com/julienschmidt/httprouter"
	"loiter/backstage/constant"
	"loiter/backstage/controller/parser"
	"loiter/backstage/controller/result"
	"loiter/backstage/controller/validator"
	"loiter/backstage/foundation"
	"loiter/backstage/service"
	"loiter/kernel/model/receiver"
	"net/http"
)

/**
 * 负载均衡控制器
 * @auth eyesYeager
 * @date 2024/1/5 16:43
 */

// UpdateAppBalancer
// @Summary			更新应用负载均衡策略
// @Description		权限：admin
// @Tags			balancer
// @Accept			json
// @Produce			json
// @Security		token
// @Param			token								header		string		true		"身份令牌"
// @Param			receiver.UpdateAppBalancer			body		string		false		"请求参数"
// @Success			200									{object}	result.Response
// @Failure			400									{object}	result.Response
// @Router			/balancer/updateAppBalancer [post]
func UpdateAppBalancer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	userClaims, err := foundation.AuthFoundation.TokenAnalysis(w, r, constant.Role.Admin)
	if err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 参数校验
	var data receiver.UpdateAppBalancer
	if err = parser.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err = validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err = service.BalancerService.UpdateAppBalancer(r, userClaims, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
