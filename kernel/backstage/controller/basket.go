package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"zliway/kernel/backstage/controller/helper"
	"zliway/kernel/backstage/controller/result"
	"zliway/kernel/backstage/controller/validator"
	"zliway/kernel/backstage/model/receiver"
	"zliway/kernel/backstage/service"
)

/**
 * basket控制器
 * @author eyesYeager
 * @date 2023/4/26 16:49
 */

// AddBasket 给应用添加组
func AddBasket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 权限校验
	if !helper.CheckAuth(r) {
		result.FailByCustom(w, r, result.Results.AuthError)
		return
	}

	// 参数校验
	var data receiver.BasketAdd
	if err := helper.PostData(r, &data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}
	if err := validator.Checker.Struct(data); err != nil {
		result.FailAttachedMsg(w, r, err.Error())
		return
	}

	// 执行业务
	if err := service.BasketService.AddBasket(r, data); err == nil {
		result.SuccessDefault(w, r, nil)
	} else {
		result.FailAttachedMsg(w, r, err.Error())
	}
}
